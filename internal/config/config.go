package config

import "time"

var COMMON = struct {
	MAX_NUM_OF_ACCOUNTS int
}{
	MAX_NUM_OF_ACCOUNTS: 6,
}

var CALCULATOR = struct {
	REINVEST_EVERY_DAYS int
}{
	REINVEST_EVERY_DAYS: 7,
}

var CMC_API = struct {
	ADDRESS         string
	FREE_LIMIT      string
	REQ_PER_MINUTES time.Duration
}{
	ADDRESS:         "https://pro-api.coinmarketcap.com/v1",
	FREE_LIMIT:      "200",
	REQ_PER_MINUTES: 5 * time.Minute,
}

var SIGNUM_API = struct {
	HOSTS                     []string
	DEFAULT_AVG_COMMIT        float64
	DEFAULT_BASE_TARGET       float64
	DEFAULT_BLOCK_REWARD      float64
	GET_AVG_COMMIT_TIME       time.Duration
	AVERAGING_DAYS_QUANTITY   uint
	CACHE_TTL                 time.Duration
	NOTIFIER_PERIOD           time.Duration
	NOTIFIER_CHECK_BLOCKS_PER uint
}{
	HOSTS: []string{
		"https://europe1.signum.network",
		"https://europe.signum.network",
		"https://canada.signum.network",
		"https://australia.signum.network",
		"https://europe2.signum.network",
		"https://europe3.signum.network",
		"https://brazil.signum.network",
		"https://uk.signum.network",
		"https://wallet.burstcoin.ro",
	},
	DEFAULT_AVG_COMMIT:        2500,
	DEFAULT_BASE_TARGET:       280000,
	DEFAULT_BLOCK_REWARD:      134,
	GET_AVG_COMMIT_TIME:       time.Hour, // per hour
	AVERAGING_DAYS_QUANTITY:   7,         // during 7 days
	CACHE_TTL:                 5 * time.Minute,
	NOTIFIER_PERIOD:           5 * time.Minute,
	NOTIFIER_CHECK_BLOCKS_PER: 6, // 5 min * 6 = per 30 min
}
