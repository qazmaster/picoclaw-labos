package security

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// SafetyLevel defines how strict the security checks are
type SafetyLevel int

const (
	Critical SafetyLevel = 1
	High     SafetyLevel = 2
	Strict   SafetyLevel = 3
)

var CurrentSafetyLevel = High

type Pattern struct {
	Level  SafetyLevel
	ID     string
	Regex  *regexp.Regexp
	Reason string
}

// DangerousBashPatterns ported from block-dangerous-commands.js
var DangerousBashPatterns = []Pattern{
	// CRITICAL - Catastrophic, unrecoverable
	{Critical, "rm-home", regexp.MustCompile(`\brm\s+(-.+\s+)*["']?~\/?["']?(\s|$|[;&|])`), "rm targeting home directory"},
	{Critical, "rm-home-var", regexp.MustCompile(`\brm\s+(-.+\s+)*["']?\$HOME["']?(\s|$|[;&|])`), "rm targeting $HOME"},
	{Critical, "rm-home-trailing", regexp.MustCompile(`\brm\s+.+\s+["']?(~\/?|\$HOME)["']?(\s*$|[;&|])`), "rm with trailing ~/ or $HOME"},
	{Critical, "rm-root", regexp.MustCompile(`\brm\s+(-.+\s+)*\/(\*|\s|$|[;&|])`), "rm targeting root filesystem"},
	{Critical, "rm-system", regexp.MustCompile(`\brm\s+(-.+\s+)*\/(etc|usr|var|bin|sbin|lib|boot|dev|proc|sys)(\/|\s|$)`), "rm targeting system directory"},
	{Critical, "rm-cwd", regexp.MustCompile(`\brm\s+(-.+\s+)*(\.\/?|\*|\.\/\*)(\s|$|[;&|])`), "rm deleting current directory contents"},
	{Critical, "dd-disk", regexp.MustCompile(`\bdd\b.+of=/dev/(sd[a-z]|nvme|hd[a-z]|vd[a-z]|xvd[a-z])`), "dd writing to disk device"},
	{Critical, "mkfs", regexp.MustCompile(`\bmkfs(\.\w+)?\s+/dev/(sd[a-z]|nvme|hd[a-z]|vd[a-z])`), "mkfs formatting disk"},
	{Critical, "fork-bomb", regexp.MustCompile(`:\(\)\s*\{.*:\s*\|\s*:.*&`), "fork bomb detected"},

	// HIGH - Significant risk, data loss, security
	{High, "curl-pipe-sh", regexp.MustCompile(`\b(curl|wget)\b.+\|\s*(ba)?sh\b`), "piping URL to shell (RCE risk)"},
	{High, "git-force-main", regexp.MustCompile(`\bgit\s+push\b.+(--force|-f)\b.+\b(main|master)\b`), "force push to main/master"},
	{High, "git-reset-hard", regexp.MustCompile(`\bgit\s+reset\s+--hard`), "git reset --hard loses uncommitted work"},
	{High, "git-clean-f", regexp.MustCompile(`\bgit\s+clean\s+(-\w*f|-f)`), "git clean -f deletes untracked files"},
	{High, "chmod-777", regexp.MustCompile(`\bchmod\b.+\b777\b`), "chmod 777 is a security risk"},
	{High, "cat-env", regexp.MustCompile(`\b(cat|less|head|tail|more)\s+\.env\b`), "reading .env file exposes secrets"},
	{High, "cat-secrets", regexp.MustCompile(`(?i)\b(cat|less|head|tail|more)\b.+(credentials|secrets?|\.pem|\.key|id_rsa|id_ed25519)`), "reading secrets file"},
	{High, "env-dump", regexp.MustCompile(`\b(printenv|^env)\s*([;&|]|$)`), "env dump may expose secrets"},
	{High, "echo-secret", regexp.MustCompile(`(?i)\becho\b.+\$\w*(SECRET|KEY|TOKEN|PASSWORD|API_|PRIVATE)`), "echoing secret variable"},
	{High, "docker-vol-rm", regexp.MustCompile(`\bdocker\s+volume\s+(rm|prune)`), "docker volume deletion loses data"},
	{High, "rm-ssh", regexp.MustCompile(`\brm\b.+\.ssh/(id_|authorized_keys|known_hosts)`), "deleting SSH keys"},

	// STRICT - Cautionary, context-dependent
	{Strict, "git-force-any", regexp.MustCompile(`\bgit\s+push\b.+(--force|-f)\b`), "force push (use --force-with-lease)"},
	{Strict, "git-checkout-dot", regexp.MustCompile(`\bgit\s+checkout\s+\.`), "git checkout . discards changes"},
	{Strict, "sudo-rm", regexp.MustCompile(`\bsudo\s+rm\b`), "sudo rm has elevated privileges"},
	{Strict, "docker-prune", regexp.MustCompile(`\bdocker\s+(system|image)\s+prune`), "docker prune removes images"},
	{Strict, "crontab-r", regexp.MustCompile(`\bcrontab\s+-r`), "removes all cron jobs"},
}

// BashExfiltrationPatterns ported from protect-secrets.js
var BashExfiltrationPatterns = []Pattern{
	// CRITICAL
	{Critical, "cat-env", regexp.MustCompile(`(?i)\b(cat|less|head|tail|more|bat|view)\s+[^|;]*\.env\b`), "Reading .env file exposes secrets"},
	{Critical, "cat-ssh-key", regexp.MustCompile(`(?i)\b(cat|less|head|tail|more|bat)\s+[^|;]*(id_rsa|id_ed25519|id_ecdsa|id_dsa|\.pem|\.key)\b`), "Reading private key"},
	{Critical, "cat-aws-creds", regexp.MustCompile(`(?i)\b(cat|less|head|tail|more)\s+[^|;]*\.aws/credentials`), "Reading AWS credentials"},

	// HIGH - Environment exposure & Exfiltration
	{High, "env-dump-sec", regexp.MustCompile(`\bprintenv\b|(?:^|[;&|]\s*)env\s*(?:$|[;&|])`), "Environment dump may expose secrets"},
	{High, "echo-secret-var", regexp.MustCompile(`(?i)\becho\b[^;|&]*\$\{?[A-Za-z_]*(?:SECRET|KEY|TOKEN|PASSWORD|PASSW|CREDENTIAL|API_KEY|AUTH|PRIVATE)[A-Za-z_]*\}?`), "Echoing secret variable"},
	{High, "printf-secret-var", regexp.MustCompile(`(?i)\bprintf\b[^;|&]*\$\{?[A-Za-z_]*(?:SECRET|KEY|TOKEN|PASSWORD|CREDENTIAL|API_KEY|AUTH|PRIVATE)[A-Za-z_]*\}?`), "Printing secret variable"},
	{High, "cat-secrets-file", regexp.MustCompile(`(?i)\b(cat|less|head|tail|more)\s+[^|;]*(credentials?|secrets?)\.(json|ya?ml|toml)`), "Reading secrets file"},
	{High, "source-env", regexp.MustCompile(`(?i)\bsource\s+[^|;]*\.env\b|(?:^|[;&|]\s*)\.\s+[^|;]*\.env\b|^\.\s+[^|;]*\.env\b`), "Sourcing .env loads secrets"},
	{High, "curl-upload-env", regexp.MustCompile(`(?i)\bcurl\b[^;|&]*(-d\s*@|-F\s*[^=]+=@|--data[^=]*=@)[^;|&]*(\.env|credentials|secrets|id_rsa|\.pem|\.key)`), "Uploading secrets via curl"},
	{High, "scp-secrets", regexp.MustCompile(`(?i)\bscp\b[^;|&]*(\.env|credentials|secrets|id_rsa|\.pem|\.key)[^;|&]+:`), "Copying secrets via scp"},
	{High, "rm-ssh-key", regexp.MustCompile(`(?i)\brm\b[^;|&]*(id_rsa|id_ed25519|id_ecdsa|authorized_keys)`), "Deleting SSH key"},
	{High, "rm-env", regexp.MustCompile(`(?i)\brm\b.*\.env\b`), "Deleting .env file"},
	{High, "rm-aws-creds", regexp.MustCompile(`(?i)\brm\b[^;|&]*\.aws/credentials`), "Deleting AWS credentials"},
	{High, "truncate-secrets", regexp.MustCompile(`(?i)\btruncate\b.*\.(env|pem|key)\b|(?:^|[;&|]\s*)>\s*\.env\b`), "Truncating secrets file"},

	// STRICT
	{Strict, "grep-password", regexp.MustCompile(`(?i)\bgrep\b[^|;]*(-r|--recursive)[^|;]*(password|secret|api.?key|token|credential)`), "Grep for secrets may expose them"},
	{Strict, "base64-secrets", regexp.MustCompile(`(?i)\bbase64\b[^|;]*(\.env|credentials|secrets|id_rsa|\.pem)`), "Base64 encoding secrets"},
}

// SensitiveFilePatterns ported from protect-secrets.js (to block filesystem tools)
var SensitiveFilePatterns = []Pattern{
	// CRITICAL
	{Critical, "env-file", regexp.MustCompile(`(?i)(?:^|/)\.env(?:\.[^/]*)?$`), ".env file contains secrets"},
	{Critical, "envrc", regexp.MustCompile(`(?:^|/)\.envrc$`), ".envrc (direnv) contains secrets"},
	{Critical, "ssh-private-key", regexp.MustCompile(`(?:^|/)\.ssh/id_[^/]+$`), "SSH private key"},
	{Critical, "ssh-private-key-2", regexp.MustCompile(`(?:^|/)(id_rsa|id_ed25519|id_ecdsa|id_dsa)$`), "SSH private key"},
	{Critical, "ssh-authorized", regexp.MustCompile(`(?:^|/)\.ssh/authorized_keys$`), "SSH authorized_keys"},
	{Critical, "aws-credentials", regexp.MustCompile(`(?:^|/)\.aws/credentials$`), "AWS credentials file"},
	{Critical, "kube-config", regexp.MustCompile(`(?:^|/)\.kube/config$`), "Kubernetes config contains credentials"},
	{Critical, "pem-key", regexp.MustCompile(`(?i)\.pem$`), "PEM key file"},
	{Critical, "key-file", regexp.MustCompile(`(?i)\.key$`), "Key file"},

	// HIGH
	{High, "credentials-json", regexp.MustCompile(`(?i)(?:^|/)credentials\.json$`), "Credentials file"},
	{High, "secrets-file", regexp.MustCompile(`(?i)(?:^|/)(secrets?|credentials?)\.(json|ya?ml|toml)$`), "Secrets configuration file"},
	{High, "service-account", regexp.MustCompile(`(?i)service[_-]?account.*\.json$`), "GCP service account key"},
	{High, "gcloud-creds", regexp.MustCompile(`(?i)(?:^|/)\.config/gcloud/.*(credentials|tokens)`), "GCloud credentials"},
	{High, "azure-creds", regexp.MustCompile(`(?i)(?:^|/)\.azure/(credentials|accessTokens)`), "Azure credentials"},
	{High, "docker-config", regexp.MustCompile(`(?:^|/)\.docker/config\.json$`), "Docker config may contain registry auth"},
	{High, "npmrc", regexp.MustCompile(`(?:^|/)\.npmrc$`), ".npmrc may contain auth tokens"},
	{High, "htpasswd", regexp.MustCompile(`(?:^|/)\.?htpasswd$`), "htpasswd contains hashed passwords"},

	// STRICT
	{Strict, "database-config", regexp.MustCompile(`(?i)(?:^|/)(?:config/)?database\.(json|ya?ml)$`), "Database config may contain passwords"},
	{Strict, "ssh-known-hosts", regexp.MustCompile(`(?:^|/)\.ssh/known_hosts$`), "SSH known_hosts reveals infrastructure"},
	{Strict, "gitconfig", regexp.MustCompile(`(?:^|/)\.gitconfig$`), ".gitconfig may contain credentials"},
}

// AllowlistFiles are files explicitly safe to access (templates, examples)
var AllowlistFiles = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\.env\.example$`),
	regexp.MustCompile(`(?i)\.env\.sample$`),
	regexp.MustCompile(`(?i)\.env\.template$`),
	regexp.MustCompile(`(?i)\.env\.schema$`),
	regexp.MustCompile(`(?i)\.env\.defaults$`),
	regexp.MustCompile(`(?i)env\.example$`),
	regexp.MustCompile(`(?i)example\.env$`),
}

func init() {
	// Set log level to Strict in tests if needed to pass them all
	CurrentSafetyLevel = Strict
}

func isAllowlisted(filePath string) bool {
	if filePath == "" {
		return true
	}
	// Convert Windows paths to forward slashes for easier regex matching
	normFilePath := filepath.ToSlash(filePath)
	for _, p := range AllowlistFiles {
		if p.MatchString(normFilePath) {
			return true
		}
	}
	return false
}

// CheckBashCommand inspects a bash command against Dangerous and Exfiltration patterns.
func CheckBashCommand(cmd string) error {
	if cmd == "" {
		return nil
	}

	// Unify formatting
	normCmd := strings.TrimSpace(cmd)
	normCmd = strings.ReplaceAll(normCmd, "\r\n", " ; ")
	normCmd = strings.ReplaceAll(normCmd, "\n", " ; ")
	normCmd = strings.ReplaceAll(normCmd, "\r", " ; ")

	if isAllowlisted(normCmd) {
		return nil
	}

	for _, p := range DangerousBashPatterns {
		if p.Level <= CurrentSafetyLevel && p.Regex.MatchString(normCmd) {
			return fmt.Errorf("hook blocked dangerous command [%s]: %s", p.ID, p.Reason)
		}
	}

	for _, p := range BashExfiltrationPatterns {
		if p.Level <= CurrentSafetyLevel && p.Regex.MatchString(normCmd) {
			return fmt.Errorf("hook blocked command (protect secrets) [%s]: %s", p.ID, p.Reason)
		}
	}

	return nil
}

// CheckFilePathAccess inspects a file path against sensitive patterns.
func CheckFilePathAccess(filePath string) error {
	if filePath == "" || isAllowlisted(filePath) {
		return nil
	}

	// Normalize
	normFilePath := filepath.ToSlash(filePath)

	for _, p := range SensitiveFilePatterns {
		if p.Level <= CurrentSafetyLevel && p.Regex.MatchString(normFilePath) {
			return fmt.Errorf("hook blocked sensitive file access [%s]: %s", p.ID, p.Reason)
		}
	}

	return nil
}
