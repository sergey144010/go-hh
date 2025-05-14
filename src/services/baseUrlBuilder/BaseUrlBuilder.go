package BaseUrlBuilder

import (
	"fmt"
)

type BaseUrlBuilder struct {
	BaseUrl string
}

func (builder *BaseUrlBuilder) Get() string {
	return builder.BaseUrl
}

func (builder *BaseUrlBuilder) Text(text string) *BaseUrlBuilder {
	builder.BaseUrl = builder.BaseUrl + "text=" + fmt.Sprintf("%s", text) + "&"
	return builder
}

func (builder *BaseUrlBuilder) Area(area int) *BaseUrlBuilder {
	builder.BaseUrl = builder.BaseUrl + "area=" + fmt.Sprintf("%d", area) + "&"
	return builder
}

func (builder *BaseUrlBuilder) Period(period int) *BaseUrlBuilder {
	builder.BaseUrl = builder.BaseUrl + "period=" + fmt.Sprintf("%d", period) + "&"
	return builder
}

func (builder *BaseUrlBuilder) ResponsesCountEnabled() *BaseUrlBuilder {
	builder.BaseUrl = builder.BaseUrl + "responses_count_enabled=true&"
	return builder
}

// "https://api.hh.ru/vacancies?text=php&area=113&period=1&responses_count_enabled=true"