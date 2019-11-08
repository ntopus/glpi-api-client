package glpi_api_client

type CreateTicket struct {
	Name         string `json:"name"`
	Content      string `json:"content"`
	Status       int    `json:"status"`
	Urgency      int    `json:"urgency"`
	Impact       int    `json:"impact"`
	DisableNotif bool   `json:"_disablenotif"`
}

type AddFollowupTicket struct {
	IsPrivate      string `json:"is_private"`
	RequestTypesId string `json:"requesttypes_id"`
	Content        string `json:"content"`
}

type ReadTicket struct {
	Id                       int    `json:"id"`
	EntitiesId               int    `json:"entities_id"`
	Name                     string `json:"name"`
	Date                     string `json:"date"`
	CloseDate                string `json:"closedate"`
	SolveDate                string `json:"solvedate"`
	DateMod                  string `json:"date_mod"`
	UsersIdLastUpdater       int    `json:"users_id_lastupdater"`
	Status                   int    `json:"status"`
	UsersIdRecipient         int    `json:"users_id_recipient"`
	RequestTypesId           int    `json:"requesttypes_id"`
	Content                  string `json:"content"`
	Urgency                  int    `json:"urgency"`
	Impact                   int    `json:"impact"`
	Priority                 int    `json:"priority"`
	ItilCategoriesId         int    `json:"itilcategories_id"`
	Type                     int    `json:"type"`
	GlobalValidation         int    `json:"global_validation"`
	SlasIdTtr                int    `json:"slas_id_ttr"`
	SlasIdTto                int    `json:"slas_id_tto"`
	SlaLevelsIdTtr           int    `json:"slalevels_id_ttr"`
	TimeToResolve            string `json:"time_to_resolve"`
	TimeToOwn                string `json:"time_to_own"`
	BeginWaitingDate         string `json:"begin_waiting_date"`
	SlaWaitingDuration       int    `json:"sla_waiting_duration"`
	OlaWaitingDuration       int    `json:"ola_waiting_duration"`
	OlasIdTto                int    `json:"olas_id_tto"`
	OlasIdTtr                int    `json:"olas_id_ttr"`
	OlaLevelsIdTtr           int    `json:"olalevels_id_ttr"`
	InternalTimeToResolve    string `json:"internal_time_to_resolve"`
	InternalTimeToOwn        string `json:"internal_time_to_own"`
	WaitingDuration          int    `json:"waiting_duration"`
	CloseDelayStat           int    `json:"close_delay_stat"`
	SolveDelayStat           int    `json:"solve_delay_stat"`
	TakeIntoAccountDelayStat int    `json:"takeintoaccount_delay_stat"`
	ActionTime               int    `json:"actiontime"`
	IsDeleted                int    `json:"is_deleted"`
	LocationsId              int    `json:"locations_id"`
	ValidationPercent        int    `json:"validation_percent"`
	DateCreation             string `json:"date_creation"`
	Links                    []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}
