package faker

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

/**
 * @link https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
 * On date of 2016-04-22
 */
var languageCode = []string{"aa", "ab", "ae", "af", "ak", "am", "an", "ar", "as", "av", "ay", "az", "ba", "be", "bg", "bh", "bi", "bm", "bn", "bo", "br", "bs", "ca", "ce", "ch", "co", "cr", "cs", "cu", "cv", "cy", "da", "de", "dv", "dz", "ee", "el", "en", "eo", "es", "et", "eu", "fa", "ff", "fi", "fj", "fo", "fr", "fy", "ga", "gd", "gl", "gn", "gu", "gv", "ha", "he", "hi", "ho", "hr", "ht", "hu", "hy", "hz", "ia", "id", "ie", "ig", "ii", "ik", "io", "is", "it", "iu", "ja", "jv", "ka", "kg", "ki", "kj", "kk", "kl", "km", "kn", "ko", "kr", "ks", "ku", "kv", "kw", "ky", "la", "lb", "lg", "li", "ln", "lo", "lt", "lu", "lv", "mg", "mh", "mi", "mk", "ml", "mn", "mr", "ms", "mt", "my", "na", "nb", "nd", "ne", "ng", "nl", "nn", "no", "nr", "nv", "ny", "oc", "oj", "om", "or", "os", "pa", "pi", "pl", "ps", "pt", "qu", "rm", "rn", "ro", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm", "sn", "so", "sq", "sr", "ss", "st", "su", "sv", "sw", "ta", "te", "tg", "th", "ti", "tk", "tl", "tn", "to", "tr", "ts", "tt", "tw", "ty", "ug", "uk", "ur", "uz", "ve", "vi", "vo", "wa", "wo", "xh", "yi", "yo", "za", "zh", "zu"}

/**
 * @link https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
 * On date of 2014-10-19
 */
var countryCode = []string{"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT", "ZA", "ZM", "ZW"}

/**
 * @link https://en.wikipedia.org/wiki/ISO_4217
 * On date of 2017-07-07
 *
 * With the following exceptions:
 * SVC has been replaced by the USD in 2001: https://en.wikipedia.org/wiki/Salvadoran_col%C3%B3n
 * ZWL has been suspended since 2009: https://en.wikipedia.org/wiki/Zimbabwean_dollar
 */
var currencyCode = []string{"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRO", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS", "SRD", "SSP", "STD", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VEF", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW"}

/**
 * @link https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
 * On date of 2014-10-19
 */
var countryISOAlpha3 = []string{"ABW", "AFG", "AGO", "AIA", "ALA", "ALB", "AND", "ARE", "ARG", "ARM", "ASM", "ATA", "ATF", "ATG", "AUS", "AUT", "AZE", "BDI", "BEL", "BEN", "BES", "BFA", "BGD", "BGR", "BHR", "BHS", "BIH", "BLM", "BLR", "BLZ", "BMU", "BOL", "BRA", "BRB", "BRN", "BTN", "BVT", "BWA", "CAF", "CAN", "CCK", "CHE", "CHL", "CHN", "CIV", "CMR", "COD", "COG", "COK", "COL", "COM", "CPV", "CRI", "CUB", "CUW", "CXR", "CYM", "CYP", "CZE", "DEU", "DJI", "DMA", "DNK", "DOM", "DZA", "ECU", "EGY", "ERI", "ESH", "ESP", "EST", "ETH", "FIN", "FJI", "FLK", "FRA", "FRO", "FSM", "GAB", "GBR", "GEO", "GGY", "GHA", "GIB", "GIN", "GLP", "GMB", "GNB", "GNQ", "GRC", "GRD", "GRL", "GTM", "GUF", "GUM", "GUY", "HKG", "HMD", "HND", "HRV", "HTI", "HUN", "IDN", "IMN", "IND", "IOT", "IRL", "IRN", "IRQ", "ISL", "ISR", "ITA", "JAM", "JEY", "JOR", "JPN", "KAZ", "KEN", "KGZ", "KHM", "KIR", "KNA", "KOR", "KWT", "LAO", "LBN", "LBR", "LBY", "LCA", "LIE", "LKA", "LSO", "LTU", "LUX", "LVA", "MAC", "MAF", "MAR", "MCO", "MDA", "MDG", "MDV", "MEX", "MHL", "MKD", "MLI", "MLT", "MMR", "MNE", "MNG", "MNP", "MOZ", "MRT", "MSR", "MTQ", "MUS", "MWI", "MYS", "MYT", "NAM", "NCL", "NER", "NFK", "NGA", "NIC", "NIU", "NLD", "NOR", "NPL", "NRU", "NZL", "OMN", "PAK", "PAN", "PCN", "PER", "PHL", "PLW", "PNG", "POL", "PRI", "PRK", "PRT", "PRY", "PSE", "PYF", "QAT", "REU", "ROU", "RUS", "RWA", "SAU", "SDN", "SEN", "SGP", "SGS", "SHN", "SJM", "SLB", "SLE", "SLV", "SMR", "SOM", "SPM", "SRB", "SSD", "STP", "SUR", "SVK", "SVN", "SWE", "SWZ", "SXM", "SYC", "SYR", "TCA", "TCD", "TGO", "THA", "TJK", "TKL", "TKM", "TLS", "TON", "TTO", "TUN", "TUR", "TUV", "TWN", "TZA", "UGA", "UKR", "UMI", "URY", "USA", "UZB", "VAT", "VCT", "VEN", "VGB", "VIR", "VNM", "VUT", "WLF", "WSM", "YEM", "ZAF", "ZMB", "ZWE"}

var locale = []string{"aa_DJ", "aa_ER", "aa_ET", "af_NA", "af_ZA", "ak_GH", "am_ET", "ar_AE", "ar_BH", "ar_DZ", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SD", "ar_SY", "ar_TN", "ar_YE", "as_IN", "az_AZ", "be_BY", "bg_BG", "bn_BD", "bn_IN", "bo_CN", "bo_IN", "bs_BA", "byn_ER", "ca_ES", "cch_NG", "cs_CZ", "cy_GB", "da_DK", "de_AT", "de_BE", "de_CH", "de_DE", "de_LI", "de_LU", "dv_MV", "dz_BT", "ee_GH", "ee_TG", "el_CY", "el_GR", "en_AS", "en_AU", "en_BE", "en_BW", "en_BZ", "en_CA", "en_GB", "en_GU", "en_HK", "en_IE", "en_IN", "en_JM", "en_MH", "en_MP", "en_MT", "en_NA", "en_NZ", "en_PH", "en_PK", "en_SG", "en_TT", "en_UM", "en_US", "en_VI", "en_ZA", "en_ZW", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_ES", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PE", "es_PR", "es_PY", "es_SV", "es_US", "es_UY", "es_VE", "et_EE", "eu_ES", "fa_AF", "fa_IR", "fi_FI", "fil_PH", "fo_FO", "fr_BE", "fr_CA", "fr_CH", "fr_FR", "fr_LU", "fr_MC", "fr_SN", "fur_IT", "ga_IE", "gaa_GH", "gez_ER", "gez_ET", "gl_ES", "gsw_CH", "gu_IN", "gv_GB", "ha_GH", "ha_NE", "ha_NG", "ha_SD", "haw_US", "he_IL", "hi_IN", "hr_HR", "hu_HU", "hy_AM", "id_ID", "ig_NG", "ii_CN", "is_IS", "it_CH", "it_IT", "ja_JP", "ka_GE", "kaj_NG", "kam_KE", "kcg_NG", "kfo_CI", "kk_KZ", "kl_GL", "km_KH", "kn_IN", "ko_KR", "kok_IN", "kpe_GN", "kpe_LR", "ku_IQ", "ku_IR", "ku_SY", "ku_TR", "kw_GB", "ky_KG", "ln_CD", "ln_CG", "lo_LA", "lt_LT", "lv_LV", "mk_MK", "ml_IN", "mn_CN", "mn_MN", "mr_IN", "ms_BN", "ms_MY", "mt_MT", "my_MM", "nb_NO", "nds_DE", "ne_IN", "ne_NP", "nl_BE", "nl_NL", "nn_NO", "nr_ZA", "nso_ZA", "ny_MW", "oc_FR", "om_ET", "om_KE", "or_IN", "pa_IN", "pa_PK", "pl_PL", "ps_AF", "pt_BR", "pt_PT", "ro_MD", "ro_RO", "ru_RU", "ru_UA", "rw_RW", "sa_IN", "se_FI", "se_NO", "sh_BA", "sh_CS", "sh_YU", "si_LK", "sid_ET", "sk_SK", "sl_SI", "so_DJ", "so_ET", "so_KE", "so_SO", "sq_AL", "sr_BA", "sr_CS", "sr_ME", "sr_RS", "sr_YU", "ss_SZ", "ss_ZA", "st_LS", "st_ZA", "sv_FI", "sv_SE", "sw_KE", "sw_TZ", "syr_SY", "ta_IN", "te_IN", "tg_TJ", "th_TH", "ti_ER", "ti_ET", "tig_ER", "tn_ZA", "to_TO", "tr_TR", "trv_TW", "ts_ZA", "tt_RU", "ug_CN", "uk_UA", "ur_IN", "ur_PK", "uz_AF", "uz_UZ", "ve_ZA", "vi_VN", "wal_ET", "wo_SN", "xh_ZA", "yo_NG", "zh_CN", "zh_HK", "zh_MO", "zh_SG", "zh_TW", "zu_ZA"}
var localeDash = []string{"aa-DJ", "aa-ER", "aa-ET", "af-NA", "af-ZA", "ak-GH", "am-ET", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IQ", "ar-JO", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-OM", "ar-QA", "ar-SA", "ar-SD", "ar-SY", "ar-TN", "ar-YE", "as-IN", "az-AZ", "be-BY", "bg-BG", "bn-BD", "bn-IN", "bo-CN", "bo-IN", "bs-BA", "byn-ER", "ca-ES", "cch-NG", "cs-CZ", "cy-GB", "da-DK", "de-AT", "de-BE", "de-CH", "de-DE", "de-LI", "de-LU", "dv-MV", "dz-BT", "ee-GH", "ee-TG", "el-CY", "el-GR", "en-AS", "en-AU", "en-BE", "en-BW", "en-BZ", "en-CA", "en-GB", "en-GU", "en-HK", "en-IE", "en-IN", "en-JM", "en-MH", "en-MP", "en-MT", "en-NA", "en-NZ", "en-PH", "en-PK", "en-SG", "en-TT", "en-UM", "en-US", "en-VI", "en-ZA", "en-ZW", "es-AR", "es-BO", "es-CL", "es-CO", "es-CR", "es-DO", "es-EC", "es-ES", "es-GT", "es-HN", "es-MX", "es-NI", "es-PA", "es-PE", "es-PR", "es-PY", "es-SV", "es-US", "es-UY", "es-VE", "et-EE", "eu-ES", "fa-AF", "fa-IR", "fi-FI", "fil-PH", "fo-FO", "fr-BE", "fr-CA", "fr-CH", "fr-FR", "fr-LU", "fr-MC", "fr-SN", "fur-IT", "ga-IE", "gaa-GH", "gez-ER", "gez-ET", "gl-ES", "gsw-CH", "gu-IN", "gv-GB", "ha-GH", "ha-NE", "ha-NG", "ha-SD", "haw-US", "he-IL", "hi-IN", "hr-HR", "hu-HU", "hy-AM", "id-ID", "ig-NG", "ii-CN", "is-IS", "it-CH", "it-IT", "ja-JP", "ka-GE", "kaj-NG", "kam-KE", "kcg-NG", "kfo-CI", "kk-KZ", "kl-GL", "km-KH", "kn-IN", "ko-KR", "kok-IN", "kpe-GN", "kpe-LR", "ku-IQ", "ku-IR", "ku-SY", "ku-TR", "kw-GB", "ky-KG", "ln-CD", "ln-CG", "lo-LA", "lt-LT", "lv-LV", "mk-MK", "ml-IN", "mn-CN", "mn-MN", "mr-IN", "ms-BN", "ms-MY", "mt-MT", "my-MM", "nb-NO", "nds-DE", "ne-IN", "ne-NP", "nl-BE", "nl-NL", "nn-NO", "nr-ZA", "nso-ZA", "ny-MW", "oc-FR", "om-ET", "om-KE", "or-IN", "pa-IN", "pa-PK", "pl-PL", "ps-AF", "pt-BR", "pt-PT", "ro-MD", "ro-RO", "ru-RU", "ru-UA", "rw-RW", "sa-IN", "se-FI", "se-NO", "sh-BA", "sh-CS", "sh-YU", "si-LK", "sid-ET", "sk-SK", "sl-SI", "so-DJ", "so-ET", "so-KE", "so-SO", "sq-AL", "sr-BA", "sr-CS", "sr-ME", "sr-RS", "sr-YU", "ss-SZ", "ss-ZA", "st-LS", "st-ZA", "sv-FI", "sv-SE", "sw-KE", "sw-TZ", "syr-SY", "ta-IN", "te-IN", "tg-TJ", "th-TH", "ti-ER", "ti-ET", "tig-ER", "tn-ZA", "to-TO", "tr-TR", "trv-TW", "ts-ZA", "tt-RU", "ug-CN", "uk-UA", "ur-IN", "ur-PK", "uz-AF", "uz-UZ", "ve-ZA", "vi-VN", "wal-ET", "wo-SN", "xh-ZA", "yo-NG", "zh-CN", "zh-HK", "zh-MO", "zh-SG", "zh-TW", "zu-ZA"}

func LanguageCode() string {
	return StringIn(languageCode)
}

func CountryCode() string {
	return StringIn(countryCode)
}

func CurrencyCode() string {
	return StringIn(currencyCode)
}

func ISOAlpha3() string {
	return StringIn(countryISOAlpha3)
}

func Local() string {
	return StringIn(locale)
}

func LocalDash() string {
	return StringIn(localeDash)
}

func Emoji() string {
	return string(RuneBetween(0x1F600, 0x1F637))
}

func Md5() string {
	h := md5.Sum([]byte(Letters(10)))
	return hex.EncodeToString(h[:])
}

func Sha1() string {
	h := sha1.Sum([]byte(Letters(10)))
	return hex.EncodeToString(h[:])
}

func Sha256() string {
	h := sha256.Sum256([]byte(Letters(10)))
	return hex.EncodeToString(h[:])
}

func Sha512() string {
	h := sha512.Sum512([]byte(Letters(10)))
	return hex.EncodeToString(h[:])
}

func Bool() bool {
	return Intn64(2) == 0
}

func BoolRatio(r int) bool {
	return int(Intn64(100)) <= r
}

