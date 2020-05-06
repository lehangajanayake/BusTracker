/**
 * Send an HTTP POST request and display the response on Serial
 */

#include <SIM808.h>
#include <ArduinoLog.h>

#if defined(__AVR__)
    #include <SoftwareSerial.h>
    #define SIM_SERIAL_TYPE	SoftwareSerial					///< Type of variable that holds the Serial communication with SIM808
    #define SIM_SERIAL		SIM_SERIAL_TYPE(SIM_TX, SIM_RX)	///< Definition of the instance that holds the Serial communication with SIM808    
    
    #define STRLCPY_P(s1, s2) strlcpy_P(s1, s2, BUFFER_SIZE)
#else
    #include <HardwareSerial.h>
    #define SIM_SERIAL_TYPE	HardwareSerial					///< Type of variable that holds the Serial communication with SIM808
    #define SIM_SERIAL		SIM_SERIAL_TYPE(2)	            ///< Definition of the instance that holds the Serial communication with SIM808    
    
    #define STRLCPY_P(s1, s2) strlcpy(s1, s2, BUFFER_SIZE)
#endif

#define SIM_RST		5	///< SIM808 RESET
#define SIM_RX		6	///< SIM808 RXD
#define SIM_TX		7	///< SIM808 TXD
#define SIM_PWR		9	///< SIM808 PWRKEY
#define SIM_STATUS	8	///< SIM808 STATUS

#define SIM808_BAUDRATE 4800    ///< Control the baudrate use to communicate with the SIM808 module
#define SERIAL_BAUDRATE 38400   ///< Controls the serial baudrate between the arduino and the computer
#define NETWORK_DELAY  10000    ///< Delay between each GPS read

#define GPRS_APN    "vodafone"  ///< Your provider Access Point Name
#define GPRS_USER   NULL        ///< Your provider APN user (usually not needed)
#define GPRS_PASS   NULL        ///< Your provider APN password (usually not needed)

#define BUFFER_SIZE 512         ///< Side of the response buffer
#define NL  "\n"

SIM_SERIAL_TYPE simSerial = SIM_SERIAL;
SIM808 sim808 = SIM808(SIM_RST, SIM_PWR, SIM_STATUS);
bool done = false;
char buffer[BUFFER_SIZE];

void setup() {
    Serial.begin(SERIAL_BAUDRATE);
    Log.begin(LOG_LEVEL_NOTICE, &Serial);

    simSerial.begin(SIM808_BAUDRATE);
    sim808.begin(simSerial);

    Log.notice(S_F("Powering on SIM808..." NL));
    sim808.powerOnOff(true);
    sim808.init();    
}

void loop() {
    if(done) {
        delay(NETWORK_DELAY);
        return;
    }

    SIM808NetworkRegistrationState status = sim808.getNetworkRegistrationStatus();
    SIM808SignalQualityReport report = sim808.getSignalQuality();

    bool isAvailable = static_cast<int8_t>(status) &
        (static_cast<int8_t>(SIM808NetworkRegistrationState::Registered) | static_cast<int8_t>(SIM808NetworkRegistrationState::Roaming))
        != 0;

    if(!isAvailable) {
        Log.notice(S_F("No network yet..." NL));
        delay(NETWORK_DELAY);
        return;
    }

    Log.notice(S_F("Network is ready." NL));
    Log.notice(S_F("Attenuation : %d dBm, Estimated quality : %d" NL), report.attenuation, report.rssi);

    bool enabled = false;
	do {
        Log.notice(S_F("Powering on SIM808's GPRS..." NL));
        enabled = sim808.enableGprs(GPRS_APN, GPRS_USER, GPRS_PASS);        
    } while(!enabled);
    
    Log.notice(S_F("Sending HTTP request..." NL));
    STRLCPY_P(buffer, PSTR("This is the body"));
    //notice that we're using the same buffer for both body and response
    uint16_t responseCode = sim808.httpPost("http://httpbin.org/anything", S_F("text/plain"), buffer, buffer, BUFFER_SIZE);

    Log.notice(S_F("Server responsed : %d" NL), responseCode);
    Log.notice(buffer);

    done = true;
}