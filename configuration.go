package smtpmock

import "fmt"

// SMTP mock configuration structure. Provides to configure mock behaviour
type configuration struct {
	hostAddress                   string
	portNumber                    int
	logToStdout                   bool
	logServerActivity             bool
	isCmdFailFast                 bool
	msgGreeting                   string
	msgInvalidCmd                 string
	msgQuitCmd                    string
	msgInvalidCmdHeloSequence     string
	msgInvalidCmdHeloArg          string
	msgHeloBlacklistedDomain      string
	msgHeloReceived               string
	msgInvalidCmdMailfromSequence string
	msgInvalidCmdMailfromArg      string
	msgMailfromBlacklistedEmail   string
	msgMailfromReceived           string
	msgInvalidCmdRcpttoSequence   string
	msgInvalidCmdRcpttoArg        string
	msgRcpttoNotRegisteredEmail   string
	msgRcpttoBlacklistedEmail     string
	msgRcpttoReceived             string
	msgInvalidCmdDataSequence     string
	msgDataReceived               string
	msgMsgSizeIsTooBig            string
	msgMsgReceived                string
	blacklistedHeloDomains        []string
	blacklistedMailfromEmails     []string
	blacklistedRcpttoEmails       []string
	notRegisteredEmails           []string
	msqSizeLimit                  int
	sessionTimeout                int

	// TODO: add ability to send 221 response before end of session for case when fail fast scenario enabled
}

// New configuration builder. Returns pointer to valid new configuration structure
func newConfiguration(config ConfigurationAttr) *configuration {
	config.assignDefaultValues()

	return &configuration{
		hostAddress:                   config.HostAddress,
		portNumber:                    config.PortNumber,
		logToStdout:                   config.LogToStdout,
		logServerActivity:             config.LogServerActivity,
		isCmdFailFast:                 config.IsCmdFailFast,
		msgGreeting:                   config.MsgGreeting,
		msgInvalidCmd:                 config.MsgInvalidCmd,
		msgInvalidCmdHeloSequence:     config.MsgInvalidCmdHeloSequence,
		msgInvalidCmdHeloArg:          config.MsgInvalidCmdHeloArg,
		msgHeloBlacklistedDomain:      config.MsgHeloBlacklistedDomain,
		msgHeloReceived:               config.MsgHeloReceived,
		msgInvalidCmdMailfromSequence: config.MsgInvalidCmdMailfromSequence,
		msgInvalidCmdMailfromArg:      config.MsgInvalidCmdMailfromArg,
		msgMailfromBlacklistedEmail:   config.MsgMailfromBlacklistedEmail,
		msgMailfromReceived:           config.MsgMailfromReceived,
		msgInvalidCmdRcpttoSequence:   config.MsgInvalidCmdRcpttoSequence,
		msgInvalidCmdRcpttoArg:        config.MsgInvalidCmdRcpttoArg,
		msgRcpttoNotRegisteredEmail:   config.MsgRcpttoNotRegisteredEmail,
		msgRcpttoBlacklistedEmail:     config.MsgRcpttoBlacklistedEmail,
		msgRcpttoReceived:             config.MsgRcpttoReceived,
		msgInvalidCmdDataSequence:     config.MsgInvalidCmdDataSequence,
		msgDataReceived:               config.MsgDataReceived,
		msgMsgSizeIsTooBig:            config.MsgMsgSizeIsTooBig,
		msgMsgReceived:                config.MsgMsgReceived,
		msgQuitCmd:                    config.MsgQuitCmd,
		blacklistedHeloDomains:        config.BlacklistedHeloDomains,
		blacklistedMailfromEmails:     config.BlacklistedMailfromEmails,
		blacklistedRcpttoEmails:       config.BlacklistedRcpttoEmails,
		notRegisteredEmails:           config.NotRegisteredEmails,
		msqSizeLimit:                  config.MsqSizeLimit,
		sessionTimeout:                config.SessionTimeout,
	}
}

// ConfigurationAttr kwargs structure for configuration builder
type ConfigurationAttr struct {
	HostAddress                   string
	PortNumber                    int
	LogToStdout                   bool
	LogServerActivity             bool
	IsCmdFailFast                 bool
	MsgGreeting                   string
	MsgInvalidCmd                 string
	MsgQuitCmd                    string
	MsgInvalidCmdHeloSequence     string
	MsgInvalidCmdHeloArg          string
	MsgHeloBlacklistedDomain      string
	MsgHeloReceived               string
	MsgInvalidCmdMailfromSequence string
	MsgInvalidCmdMailfromArg      string
	MsgMailfromBlacklistedEmail   string
	MsgMailfromReceived           string
	MsgInvalidCmdRcpttoSequence   string
	MsgInvalidCmdRcpttoArg        string
	MsgRcpttoNotRegisteredEmail   string
	MsgRcpttoBlacklistedEmail     string
	MsgRcpttoReceived             string
	MsgInvalidCmdDataSequence     string
	MsgDataReceived               string
	MsgMsgSizeIsTooBig            string
	MsgMsgReceived                string
	BlacklistedHeloDomains        []string
	BlacklistedMailfromEmails     []string
	BlacklistedRcpttoEmails       []string
	NotRegisteredEmails           []string
	MsqSizeLimit                  int
	SessionTimeout                int
}

// ConfigurationAttr methods

// Assigns server defaults
func (config *ConfigurationAttr) assignServerDefaultValues() {
	if config.HostAddress == emptyString {
		config.HostAddress = defaultHostAddress
	}
	if config.MsgGreeting == emptyString {
		config.MsgGreeting = defaultGreetingMsg
	}
	if config.MsgInvalidCmd == emptyString {
		config.MsgInvalidCmd = defaultInvalidCmdMsg
	}
	if config.MsgQuitCmd == emptyString {
		config.MsgQuitCmd = defaultQuitMsg
	}
	if config.SessionTimeout == 0 {
		config.SessionTimeout = defaultSessionTimeout
	}
}

// Assigns handlerHelo defaults
func (config *ConfigurationAttr) assignHandlerHeloDefaultValues() {
	if config.MsgInvalidCmdHeloSequence == emptyString {
		config.MsgInvalidCmdHeloSequence = defaultInvalidCmdHeloSequenceMsg
	}
	if config.MsgInvalidCmdHeloArg == emptyString {
		config.MsgInvalidCmdHeloArg = defaultInvalidCmdHeloArgMsg
	}
	if config.MsgHeloBlacklistedDomain == emptyString {
		config.MsgHeloBlacklistedDomain = defaultQuitMsg
	}
	if config.MsgHeloReceived == emptyString {
		config.MsgHeloReceived = defaultReceivedMsg
	}
}

// Assigns handlerMailfrom defaults
func (config *ConfigurationAttr) assignHandlerMailfromDefaultValues() {
	if config.MsgInvalidCmdMailfromSequence == emptyString {
		config.MsgInvalidCmdMailfromSequence = defaultInvalidCmdMailfromSequenceMsg
	}
	if config.MsgInvalidCmdMailfromArg == emptyString {
		config.MsgInvalidCmdMailfromArg = defaultInvalidCmdMailfromArgMsg
	}
	if config.MsgMailfromBlacklistedEmail == emptyString {
		config.MsgMailfromBlacklistedEmail = defaultQuitMsg
	}
	if config.MsgMailfromReceived == emptyString {
		config.MsgMailfromReceived = defaultReceivedMsg
	}
}

// Assigns handlerRcptto defaults
func (config *ConfigurationAttr) assignHandlerRcpttoDefaultValues() {
	if config.MsgInvalidCmdRcpttoSequence == emptyString {
		config.MsgInvalidCmdRcpttoSequence = defaultInvalidCmdRcpttoSequenceMsg
	}
	if config.MsgInvalidCmdRcpttoArg == emptyString {
		config.MsgInvalidCmdRcpttoArg = defaultInvalidCmdRcpttoArgMsg
	}
	if config.MsgRcpttoBlacklistedEmail == emptyString {
		config.MsgRcpttoBlacklistedEmail = defaultQuitMsg
	}
	if config.MsgRcpttoNotRegisteredEmail == emptyString {
		config.MsgRcpttoNotRegisteredEmail = defaultNotRegistredRcpttoEmailMsg
	}
	if config.MsgRcpttoReceived == emptyString {
		config.MsgRcpttoReceived = defaultReceivedMsg
	}
}

// Assigns handlerData defaults
func (config *ConfigurationAttr) assignHandlerDataDefaultValues() {
	if config.MsgInvalidCmdDataSequence == emptyString {
		config.MsgInvalidCmdDataSequence = defaultInvalidCmdDataSequenceMsg
	}
	if config.MsgDataReceived == emptyString {
		config.MsgDataReceived = defaultReadyForReceiveMsg
	}
}

// Assigns handlerMessage defaults
func (config *ConfigurationAttr) assignHandlerMessageDefaultValues() {
	if config.MsqSizeLimit == 0 {
		config.MsqSizeLimit = defaultMessageSizeLimit
	}
	if config.MsgMsgSizeIsTooBig == emptyString {
		config.MsgMsgSizeIsTooBig = fmt.Sprintf(defaultMsgSizeIsTooBigMsg+" %d bytes", config.MsqSizeLimit)
	}
	if config.MsgMsgReceived == emptyString {
		config.MsgMsgReceived = defaultReceivedMsg
	}
}

// Assigns default values to ConfigurationAttr fields
func (config *ConfigurationAttr) assignDefaultValues() {
	config.assignServerDefaultValues()
	config.assignHandlerHeloDefaultValues()
	config.assignHandlerMailfromDefaultValues()
	config.assignHandlerRcpttoDefaultValues()
	config.assignHandlerDataDefaultValues()
	config.assignHandlerMessageDefaultValues()
}