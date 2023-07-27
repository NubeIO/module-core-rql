# module-core-rql

port from https://github.com/NubeIO/module-core-rql


```bash
go build -o module-core-rql
```
to build and run rubix-os you can use the bash script
```bash
bash build.bash <YOUR_ROS_PATH>
```
example
```bash
bash build.bash code/go
```



# PDF

to add an image using a URL will not work the image needs to be on host

```
![](/data/module-core-rql/nube.png "Nube logo")
```


#### type Alert

```go
type Alert struct {
        Result *model.Alert `json:"result"`
        Error  string       `json:"error"`
}
```


#### type Alerts

```go
type Alerts struct {
        Result []model.Alert `json:"result"`
        Error  string        `json:"error"`
}
```


#### type Client

```go
type Client struct {
        Return    interface{}      `json:"return"`
        Err       string           `json:"err"`
        TimeTaken string           `json:"time_taken"`
        Storage   storage.IStorage `json:"-"`
}
```


#### func (*Client) AddAlert

```go
func (inst *Client) AddAlert(hostIDName string, body any) *Alert
```

#### func (*Client) Date

```go
func (inst *Client) Date() string
```

#### func (*Client) Day

```go
func (inst *Client) Day() string
```

#### func (*Client) GetAlerts

```go
func (inst *Client) GetAlerts(hostIDName string) *Alerts
```

#### func (*Client) GetCurrentWeather

```go
func (inst *Client) GetCurrentWeather(apiKey, city string) *CurrentWeatherResponse
```

#### func (*Client) GetForecast

```go
func (inst *Client) GetForecast(apiKey, city string, days int) *ForecastResponse
```

#### func (*Client) GetHosts

```go
func (inst *Client) GetHosts() *Hosts
```

#### func (*Client) GetPoint

```go
func (inst *Client) GetPoint(hostIDName, uuid string) *Point
```

#### func (*Client) GetPointHistories

```go
func (inst *Client) GetPointHistories(hostIDName string, pointUUIDs []string) *Histories
```

#### func (*Client) GetPoints

```go
func (inst *Client) GetPoints(hostIDName string) *Points
```

#### func (*Client) GetPublicHolidays

```go
func (inst *Client) GetPublicHolidays(year, countryCode string) *PublicHolidaysResponse
```

#### func (*Client) GetPublicHolidaysByState

```go
func (inst *Client) GetPublicHolidaysByState(year, countryCode, state string) *PublicHolidaysResponse
```

#### func (*Client) GetScripts

```go
func (inst *Client) GetScripts() *ScriptsResponse
```

#### func (*Client) GetVariable

```go
func (inst *Client) GetVariable(uuidName string) *VarResponse
```

#### func (*Client) GetVariables

```go
func (inst *Client) GetVariables() *VarsResponse
```

#### func (*Client) HTTPGet

```go
func (inst *Client) HTTPGet(body *HTTPBody) *HTTPGet
```

#### func (*Client) IsPublicHoliday

```go
func (inst *Client) IsPublicHoliday(year, countryCode, date string) *IsPublicHolidayResponse
```

#### func (*Client) JsonToDF

```go
func (inst *Client) JsonToDF(data any) dataframe.DataFrame
```

#### func (*Client) LimitToRange

```go
func (inst *Client) LimitToRange(value float64, range1 float64, range2 float64) float64
```
LimitToRange returns the input value clamped within the specified range

#### func (*Client) PDF

```go
func (inst *Client) PDF(pdfBody *PdfBody) *PingResponse
```

#### func (*Client) Ping

```go
func (inst *Client) Ping(ipList []string) *PingResponse
```

#### func (*Client) Print

```go
func (inst *Client) Print(x interface{})
```

#### func (*Client) PrintMany

```go
func (inst *Client) PrintMany(x ...interface{})
```

#### func (*Client) RandFloat

```go
func (inst *Client) RandFloat(range1, range2 float64) float64
```
RandFloat returns a random float64 within the specified range.

#### func (*Client) RandInt

```go
func (inst *Client) RandInt(range1, range2 int) int
```
RandInt returns a random int within the specified range.

#### func (*Client) RoundTo

```go
func (inst *Client) RoundTo(value float64, decimals uint32) float64
```
RoundTo returns the input value rounded to the specified number of decimal
places.

#### func (*Client) Scale

```go
func (inst *Client) Scale(value, inMin, inMax, outMin, outMax float64) float64
```
Scale returns the (float64) input value (between inputMin and inputMax) scaled
to a value between outputMin and outputMax

#### func (*Client) SendEmail

```go
func (inst *Client) SendEmail(body *Mail) error
```
SendEmail example let body = {

    to: ["a@nube-io.com"],
    subject: "test",
    message: "testing",
    senderAddress: "aa@nube-io.com",
    password: "abc",

};

RQL.SendEmail(body);

#### func (*Client) Slack

```go
func (inst *Client) Slack(body any)
```

#### func (*Client) Sleep

```go
func (inst *Client) Sleep(duration int)
```
Sleep will delay the program for the `duration` passed in (duration is units
seconds)

#### func (*Client) Tags

```go
func (inst *Client) Tags(tag ...string)
```

#### func (*Client) Time

```go
func (inst *Client) Time() string
```

#### func (*Client) TimeDate

```go
func (inst *Client) TimeDate() string
```

#### func (*Client) TimeDateDay

```go
func (inst *Client) TimeDateDay() string
```

#### func (*Client) TimeDateFormat

```go
func (inst *Client) TimeDateFormat(format string) string
```

#### func (*Client) TimeUTC

```go
func (inst *Client) TimeUTC() time.Time
```

#### func (*Client) TimeWithMS

```go
func (inst *Client) TimeWithMS() string
```

#### func (*Client) ToString

```go
func (inst *Client) ToString(x interface{}) string
```

#### func (*Client) VarParseArray

```go
func (inst *Client) VarParseArray(uuidName string) interface{}
```
VarParseArray [1, 2, "hello"]

let data = JSON.parse(RQL.VarParseArray("array")); RQL.Return = data;

#### func (*Client) VarParseNumber

```go
func (inst *Client) VarParseNumber(uuidName string) float64
```

#### func (*Client) VarParseObject

```go
func (inst *Client) VarParseObject(uuidName string) interface{}
```
VarParseObject `{"sp":22.3,"db":99.9}`

let data = RQL.VarParseObject("obj"); let sp = JSON.parse(data); RQL.Return =
sp["sp"];

#### func (*Client) VarParseString

```go
func (inst *Client) VarParseString(uuidName string) string
```

#### func (*Client) WeatherByTownAU

```go
func (inst *Client) WeatherByTownAU(town, state string) *WeatherByTownAUResp
```

#### func (*Client) WritePointValue

```go
func (inst *Client) WritePointValue(hostIDName, uuid string, value *model.Priority) *Point
```

#### func (*Client) WritePointValuePriority

```go
func (inst *Client) WritePointValuePriority(hostIDName, uuid string, pri int, value float64) *Point
```

#### func (*Client) Year

```go
func (inst *Client) Year() string
```

#### type CurrentWeather

```go
type CurrentWeather struct {
        Location struct {
                Name           string  `json:"name"`
                Region         string  `json:"region"`
                Country        string  `json:"country"`
                Lat            float64 `json:"lat"`
                Lon            float64 `json:"lon"`
                TzId           string  `json:"tz_id"`
                LocaltimeEpoch int     `json:"localtime_epoch"`
                Localtime      string  `json:"localtime"`
        } `json:"location"`
        Current struct {
                LastUpdatedEpoch int     `json:"last_updated_epoch"`
                LastUpdated      string  `json:"last_updated"`
                TempC            float64 `json:"temp_c"`
                TempF            float64 `json:"temp_f"`
                IsDay            int     `json:"is_day"`
                Condition        struct {
                        Text string `json:"text"`
                        Icon string `json:"icon"`
                        Code int    `json:"code"`
                } `json:"condition"`
                WindMph    float64 `json:"wind_mph"`
                WindKph    float64 `json:"wind_kph"`
                WindDegree int     `json:"wind_degree"`
                WindDir    string  `json:"wind_dir"`
                PressureMb float64 `json:"pressure_mb"`
                PressureIn float64 `json:"pressure_in"`
                PrecipMm   float64 `json:"precip_mm"`
                PrecipIn   float64 `json:"precip_in"`
                Humidity   int     `json:"humidity"`
                Cloud      int     `json:"cloud"`
                FeelslikeC float64 `json:"feelslike_c"`
                FeelslikeF float64 `json:"feelslike_f"`
                VisKm      float64 `json:"vis_km"`
                VisMiles   float64 `json:"vis_miles"`
                Uv         float64 `json:"uv"`
                GustMph    float64 `json:"gust_mph"`
                GustKph    float64 `json:"gust_kph"`
        } `json:"current"`
}
```


#### type CurrentWeatherResponse

```go
type CurrentWeatherResponse struct {
        Result *CurrentWeather
        Error  string
}
```


#### type Forecast

```go
type Forecast struct {
        Location struct {
                Name           string  `json:"name"`
                Region         string  `json:"region"`
                Country        string  `json:"country"`
                Lat            float64 `json:"lat"`
                Lon            float64 `json:"lon"`
                TzId           string  `json:"tz_id"`
                LocaltimeEpoch int     `json:"localtime_epoch"`
                Localtime      string  `json:"localtime"`
        } `json:"location"`
        Current struct {
                LastUpdatedEpoch int     `json:"last_updated_epoch"`
                LastUpdated      string  `json:"last_updated"`
                TempC            float64 `json:"temp_c"`
                TempF            float64 `json:"temp_f"`
                IsDay            int     `json:"is_day"`
                Condition        struct {
                        Text string `json:"text"`
                        Icon string `json:"icon"`
                        Code int    `json:"code"`
                } `json:"condition"`
                WindMph    float64 `json:"wind_mph"`
                WindKph    float64 `json:"wind_kph"`
                WindDegree int     `json:"wind_degree"`
                WindDir    string  `json:"wind_dir"`
                PressureMb float64 `json:"pressure_mb"`
                PressureIn float64 `json:"pressure_in"`
                PrecipMm   float64 `json:"precip_mm"`
                PrecipIn   float64 `json:"precip_in"`
                Humidity   int     `json:"humidity"`
                Cloud      int     `json:"cloud"`
                FeelslikeC float64 `json:"feelslike_c"`
                FeelslikeF float64 `json:"feelslike_f"`
                VisKm      float64 `json:"vis_km"`
                VisMiles   float64 `json:"vis_miles"`
                Uv         float64 `json:"uv"`
                GustMph    float64 `json:"gust_mph"`
                GustKph    float64 `json:"gust_kph"`
        } `json:"current"`
        Forecast struct {
                Forecastday []struct {
                        Date      string `json:"date"`
                        DateEpoch int    `json:"date_epoch"`
                        Day       struct {
                                MaxtempC          float64 `json:"maxtemp_c"`
                                MaxtempF          float64 `json:"maxtemp_f"`
                                MintempC          float64 `json:"mintemp_c"`
                                MintempF          float64 `json:"mintemp_f"`
                                AvgtempC          float64 `json:"avgtemp_c"`
                                AvgtempF          float64 `json:"avgtemp_f"`
                                MaxwindMph        float64 `json:"maxwind_mph"`
                                MaxwindKph        float64 `json:"maxwind_kph"`
                                TotalprecipMm     float64 `json:"totalprecip_mm"`
                                TotalprecipIn     float64 `json:"totalprecip_in"`
                                TotalsnowCm       float64 `json:"totalsnow_cm"`
                                AvgvisKm          float64 `json:"avgvis_km"`
                                AvgvisMiles       float64 `json:"avgvis_miles"`
                                Avghumidity       float64 `json:"avghumidity"`
                                DailyWillItRain   int     `json:"daily_will_it_rain"`
                                DailyChanceOfRain int     `json:"daily_chance_of_rain"`
                                DailyWillItSnow   int     `json:"daily_will_it_snow"`
                                DailyChanceOfSnow int     `json:"daily_chance_of_snow"`
                                Condition         struct {
                                        Text string `json:"text"`
                                        Icon string `json:"icon"`
                                        Code int    `json:"code"`
                                } `json:"condition"`
                                Uv float64 `json:"uv"`
                        } `json:"day"`
                        Astro struct {
                                Sunrise          string `json:"sunrise"`
                                Sunset           string `json:"sunset"`
                                Moonrise         string `json:"moonrise"`
                                Moonset          string `json:"moonset"`
                                MoonPhase        string `json:"moon_phase"`
                                MoonIllumination string `json:"moon_illumination"`
                                IsMoonUp         int    `json:"is_moon_up"`
                                IsSunUp          int    `json:"is_sun_up"`
                        } `json:"astro"`
                        Hour []struct {
                                TimeEpoch int     `json:"time_epoch"`
                                Time      string  `json:"time"`
                                TempC     float64 `json:"temp_c"`
                                TempF     float64 `json:"temp_f"`
                                IsDay     int     `json:"is_day"`
                                Condition struct {
                                        Text string `json:"text"`
                                        Icon string `json:"icon"`
                                        Code int    `json:"code"`
                                } `json:"condition"`
                                WindMph      float64 `json:"wind_mph"`
                                WindKph      float64 `json:"wind_kph"`
                                WindDegree   int     `json:"wind_degree"`
                                WindDir      string  `json:"wind_dir"`
                                PressureMb   float64 `json:"pressure_mb"`
                                PressureIn   float64 `json:"pressure_in"`
                                PrecipMm     float64 `json:"precip_mm"`
                                PrecipIn     float64 `json:"precip_in"`
                                Humidity     int     `json:"humidity"`
                                Cloud        int     `json:"cloud"`
                                FeelslikeC   float64 `json:"feelslike_c"`
                                FeelslikeF   float64 `json:"feelslike_f"`
                                WindchillC   float64 `json:"windchill_c"`
                                WindchillF   float64 `json:"windchill_f"`
                                HeatindexC   float64 `json:"heatindex_c"`
                                HeatindexF   float64 `json:"heatindex_f"`
                                DewpointC    float64 `json:"dewpoint_c"`
                                DewpointF    float64 `json:"dewpoint_f"`
                                WillItRain   int     `json:"will_it_rain"`
                                ChanceOfRain int     `json:"chance_of_rain"`
                                WillItSnow   int     `json:"will_it_snow"`
                                ChanceOfSnow int     `json:"chance_of_snow"`
                                VisKm        float64 `json:"vis_km"`
                                VisMiles     float64 `json:"vis_miles"`
                                GustMph      float64 `json:"gust_mph"`
                                GustKph      float64 `json:"gust_kph"`
                                Uv           float64 `json:"uv"`
                        } `json:"hour"`
                } `json:"forecastday"`
        } `json:"forecast"`
        Alerts struct {
                Alert []struct {
                        Headline    string    `json:"headline"`
                        Msgtype     string    `json:"msgtype"`
                        Severity    string    `json:"severity"`
                        Urgency     string    `json:"urgency"`
                        Areas       string    `json:"areas"`
                        Category    string    `json:"category"`
                        Certainty   string    `json:"certainty"`
                        Event       string    `json:"event"`
                        Note        string    `json:"note"`
                        Effective   time.Time `json:"effective"`
                        Expires     time.Time `json:"expires"`
                        Desc        string    `json:"desc"`
                        Instruction string    `json:"instruction"`
                } `json:"alert"`
        } `json:"alerts"`
}
```


#### type ForecastResponse

```go
type ForecastResponse struct {
        Result *Forecast
        Error  string
}
```


#### type HTTPBody

```go
type HTTPBody struct {
        Url          string `json:"url"`
        Method       string
        ResponseType string `json:"response_type"` //json, string
        Headers      map[string]string
}
```


#### type HTTPGet

```go
type HTTPGet struct {
        Result any
        Error  string
}
```


#### type Histories

```go
type Histories struct {
        Result []model.PointHistory `json:"result"`
        Error  string               `json:"error"`
}
```


#### type Host

```go
type Host struct {
        Result *model.Host `json:"result"`
        Error  string      `json:"error"`
}
```


#### type Hosts

```go
type Hosts struct {
        Result []model.Host `json:"result"`
        Error  string       `json:"error"`
}
```


#### type IsPublicHoliday

```go
type IsPublicHoliday struct {
        IsPublicHoliday bool
        Name            string
        Locations       []string `json:"locations"`
}
```


#### type IsPublicHolidayResponse

```go
type IsPublicHolidayResponse struct {
        Result *IsPublicHoliday
        Error  string
}
```


#### type Mail

```go
type Mail struct {
        To            []string
        Cc            []string
        Bcc           []string
        Subject       string
        Message       string
        SenderAddress string
        Password      string
}
```


#### type PDFResponse

```go
type PDFResponse struct {
        Result []pingResult
        Error  string
}
```


#### type PdfBody

```go
type PdfBody struct {
        Text string `json:"text"`
}
```


#### type PingResponse

```go
type PingResponse struct {
        Result any
        Error  string
}
```


#### type Point

```go
type Point struct {
        Result *model.Point `json:"result"`
        Error  string       `json:"error"`
}
```


#### type Points

```go
type Points struct {
        Result []model.Point `json:"result"`
        Error  string        `json:"error"`
}
```


#### type PublicHolidays

```go
type PublicHolidays struct {
        Date        string   `json:"date"`
        LocalName   string   `json:"localName"`
        Name        string   `json:"name"`
        CountryCode string   `json:"countryCode"`
        Fixed       bool     `json:"fixed"`
        Global      bool     `json:"global"`
        Counties    []string `json:"counties"`
        LaunchYear  int      `json:"launchYear"`
        Types       []string `json:"types"`
}
```


#### type PublicHolidaysResponse

```go
type PublicHolidaysResponse struct {
        Result []PublicHolidays
        Error  string
}
```


#### type ScriptsResponse

```go
type ScriptsResponse struct {
        Result []storage.RQLRule
        Error  string
}
```


#### type VarResponse

```go
type VarResponse struct {
        Result *storage.RQLVariables
        Error  string
}
```


#### type VarsResponse

```go
type VarsResponse struct {
        Result []storage.RQLVariables
        Error  string
}
```


#### type WeatherByTownAUResp

```go
type WeatherByTownAUResp struct {
        Result *bom.Observations
        Error  string
}
```
