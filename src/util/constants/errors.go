package constants

// Error messages used throughout the app
const (
	CorruptedConfigFileError    string = "Unable to read the config file, generating a new one from scratch..."
	EndAfterBeginDateError      string = "End date provided must be after start date provided"
	FlagsUnparsedError          string = "Flags have not been parsed"
	InvalidCommandError         string = "Invalid sub-command: %s"
	InvalidDateError            string = "Unable to parse date provided (format: \"dd-mm-yy\")"
	InvalidLoginError           string = "Unable to login, check your RE on config and the plataform!"
	MissingDateError            string = "Date argument must be provided (format: \"dd-mm-yy\")"
	MissingDateFlagError        string = "Flag --date must be provided with a valid past date (format: \"dd-mm-yy\")"
	MissingDatesError           string = "Start and finish date arguments must be provided (format: \"dd-mm-yy\")"
	MissingGardFlagError        string = "Flags --on-gard AND --off-gard must be used together with valid past dates (format: \"dd-mm-yy\")"
	MissingTimeFlagError        string = "Flag --time must be provided with a valid time (format: \"HH:mm\")"
	NewConfigFileGeneratedError string = "Done! Remember to set your employee registration number next: meliponto config --re {your_re}"
	NoCommandError              string = "You must provide a sub-command"
	NoHomeFolderError           string = "User has no home folder setup, required for meliponto"
	PastDateError               string = "Date provided must be in the past"
	PastDateTimeError           string = "Date and Time provided must be in the past"
)