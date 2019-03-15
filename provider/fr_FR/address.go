package fr_FR

import (
	"strings"

	"github.com/gol4ng/faker"
)

type Departement struct {
	Number string
	Name   string
}

func (d *Departement) String() string {
	return d.Number + " " + d.Name
}

var citySuffix = []string{"Ville", "Bourg", "-les-Bains", "-sur-Mer", "-la-Forêt", "boeuf", "nec", "dan"}
var streetPrefix = []string{"rue", "rue", "chemin", "avenue", "boulevard", "place", "impasse"}
var cityFormats = []string{
	"{{lastName}}",
	"{{lastName}}",
	"{{lastName}}",
	"{{lastName}}",
	"{{lastName}}{{citySuffix}}",
	"{{lastName}}{{citySuffix}}",
	"{{lastName}}{{citySuffix}}",
	"{{lastName}}-sur-{{lastName}}",
}
var streetNameFormats = []string{
	"{{streetPrefix}} {{lastName}}",
	"{{streetPrefix}} {{firstName}} {{lastName}}",
	"{{streetPrefix}} de {{lastName}}",
}
var streetAddressFormats = []string{
	"{{streetName}}",
	"{{buildingNumber}}, {{streetName}}",
	"{{buildingNumber}}, {{streetName}}",
	"{{buildingNumber}}, {{streetName}}",
	"{{buildingNumber}}, {{streetName}}",
	"{{buildingNumber}}, {{streetName}}",
}

var addressFormats = []string{
	"{{streetAddress}}\n{{postcode}} {{city}}",
}

var buildingNumber = []string{"%", "%#", "%#", "%#", "%##"}
var postcode = []string{"#####", "## ###"}

var country = []string{"Afghanistan", "Afrique du sud", "Albanie", "Algérie", "Allemagne", "Andorre", "Angola", "Anguilla", "Antarctique", "Antigua et Barbuda", "Antilles néerlandaises", "Arabie saoudite", "Argentine", "Arménie", "Aruba", "Australie", "Autriche", "Azerbaïdjan", "Bahamas", "Bahrain", "Bangladesh", "Belgique", "Belize", "Benin", "Bermudes (Les)", "Bhoutan", "Biélorussie", "Bolivie", "Bosnie-Herzégovine", "Botswana", "Bouvet (Îles)", "Brunei", "Brésil", "Bulgarie", "Burkina Faso", "Burundi", "Cambodge", "Cameroun", "Canada", "Cap Vert", "Cayman (Îles)", "Chili", "Chine (Rép. pop.)", "Christmas (Île)", "Chypre", "Cocos (Îles)", "Colombie", "Comores", "Cook (Îles)", "Corée du Nord", "Corée, Sud", "Costa Rica", "Croatie", "Cuba", "Côte d'Ivoire", "Danemark", "Djibouti", "Dominique", "Égypte", "El Salvador", "Émirats arabes unis", "Équateur", "Érythrée", "Espagne", "Estonie", "États-Unis", "Ethiopie", "Falkland (Île)", "Fidji (République des)", "Finlande", "France", "Féroé (Îles)", "Gabon", "Gambie", "Ghana", "Gibraltar", "Grenade", "Groenland", "Grèce", "Guadeloupe", "Guam", "Guatemala", "Guinée", "Guinée Equatoriale", "Guinée-Bissau", "Guyane", "Guyane française", "Géorgie", "Géorgie du Sud et Sandwich du Sud (Îles)", "Haïti", "Heard et McDonald (Îles)", "Honduras", "Hong Kong", "Hongrie", "Îles Mineures Éloignées des États-Unis", "Inde", "Indonésie", "Irak", "Iran", "Irlande", "Islande", "Israël", "Italie", "Jamaïque", "Japon", "Jordanie", "Kazakhstan", "Kenya", "Kirghizistan", "Kiribati", "Koweit", "La Barbad", "Laos", "Lesotho", "Lettonie", "Liban", "Libye", "Libéria", "Liechtenstein", "Lithuanie", "Luxembourg", "Macau", "Macédoine", "Madagascar", "Malaisie", "Malawi", "Maldives (Îles)", "Mali", "Malte", "Mariannes du Nord (Îles)", "Maroc", "Marshall (Îles)", "Martinique", "Maurice", "Mauritanie", "Mayotte", "Mexique", "Micronésie (États fédérés de)", "Moldavie", "Monaco", "Mongolie", "Montserrat", "Mozambique", "Myanmar", "Namibie", "Nauru", "Nepal", "Nicaragua", "Niger", "Nigeria", "Niue", "Norfolk (Îles)", "Norvège", "Nouvelle Calédonie", "Nouvelle-Zélande", "Oman", "Ouganda", "Ouzbékistan", "Pakistan", "Palau", "Panama", "Papouasie-Nouvelle-Guinée", "Paraguay", "Pays-Bas", "Philippines", "Pitcairn (Îles)", "Pologne", "Polynésie française", "Porto Rico", "Portugal", "Pérou", "Qatar", "Roumanie", "Royaume-Uni", "Russie", "Rwanda", "Rép. Dém. du Congo", "République centrafricaine", "République Dominicaine", "République tchèque", "Réunion (La)", "Sahara Occidental", "Saint Pierre et Miquelon", "Saint Vincent et les Grenadines", "Saint-Kitts et Nevis", "Saint-Marin (Rép. de)", "Sainte Hélène", "Sainte Lucie", "Samoa", "Samoa", "Seychelles", "Sierra Leone", "Singapour", "Slovaquie", "Slovénie", "Somalie", "Soudan", "Sri Lanka", "Suisse", "Suriname", "Suède", "Svalbard et Jan Mayen (Îles)", "Swaziland", "Syrie", "São Tomé et Príncipe (Rép.)", "Sénégal", "Tadjikistan", "Taiwan", "Tanzanie", "Tchad", "Territoire britannique de l'océan Indien", "Territoires français du sud", "Thailande", "Timor", "Togo", "Tokelau", "Tonga", "Trinité et Tobago", "Tunisie", "Turkménistan", "Turks et Caïques (Îles)", "Turquie", "Tuvalu", "Ukraine", "Uruguay", "Vanuatu", "Vatican (Etat du)", "Venezuela", "Vierges (Îles)", "Vierges britanniques (Îles)", "Vietnam", "Wallis et Futuna (Îles)", "Yemen", "Yougoslavie", "Zambie", "Zaïre", "Zimbabwe"}

var regions = []string{"Alsace", "Aquitaine", "Auvergne", "Bourgogne", "Bretagne", "Centre", "Champagne-Ardenne", "Corse", "Franche-Comté", "Île-de-France", "Languedoc-Roussillon", "Limousin", "Lorraine", "Midi-Pyrénées", "Nord-Pas-de-Calais", "Basse-Normandie", "Haute-Normandie", "Pays-de-Loire", "Picardie", "Poitou-Charentes", "Provence-Alpes-Côte d'Azur", "Rhone-Alpes", "Guadeloupe", "Martinique", "Guyane", "Réunion", "Saint-Pierre-et-Miquelon", "Mayotte", "Saint-Barthélémy", "Saint-Martin", "Wallis-et-Futuna", "Polynésie française", "Nouvelle-Calédonie"}

var departments = []*Departement{{Number: "01", Name: "Ain"}, {Number: "02", Name: "Aisne"}, {Number: "03", Name: "Allier"}, {Number: "04", Name: "Alpes-de-Haute-Provence"}, {Number: "05", Name: "Hautes-Alpes"}, {Number: "06", Name: "Alpes-Maritimes"}, {Number: "07", Name: "Ardèche"}, {Number: "08", Name: "Ardennes"}, {Number: "09", Name: "Ariège"}, {Number: "10", Name: "Aube"}, {Number: "11", Name: "Aude"}, {Number: "12", Name: "Aveyron"}, {Number: "13", Name: "Bouches-du-Rhône"}, {Number: "14", Name: "Calvados"}, {Number: "15", Name: "Cantal"}, {Number: "16", Name: "Charente"}, {Number: "17", Name: "Charente-Maritime"}, {Number: "18", Name: "Cher"}, {Number: "19", Name: "Corrèze"}, {Number: "2A", Name: "Corse-du-Sud"}, {Number: "2B", Name: "Haute-Corse"}, {Number: "21", Name: "Côte-d'Or"}, {Number: "22", Name: "Côtes-d'Armor"}, {Number: "23", Name: "Creuse"}, {Number: "24", Name: "Dordogne"}, {Number: "25", Name: "Doubs"}, {Number: "26", Name: "Drôme"}, {Number: "27", Name: "Eure"}, {Number: "28", Name: "Eure-et-Loir"}, {Number: "29", Name: "Finistère"}, {Number: "30", Name: "Gard"}, {Number: "31", Name: "Haute-Garonne"}, {Number: "32", Name: "Gers"}, {Number: "33", Name: "Gironde"}, {Number: "34", Name: "Hérault"}, {Number: "35", Name: "Ille-et-Vilaine"}, {Number: "36", Name: "Indre"}, {Number: "37", Name: "Indre-et-Loire"}, {Number: "38", Name: "Isère"}, {Number: "39", Name: "Jura"}, {Number: "40", Name: "Landes"}, {Number: "41", Name: "Loir-et-Cher"}, {Number: "42", Name: "Loire"}, {Number: "43", Name: "Haute-Loire"}, {Number: "44", Name: "Loire-Atlantique"}, {Number: "45", Name: "Loiret"}, {Number: "46", Name: "Lot"}, {Number: "47", Name: "Lot-et-Garonne"}, {Number: "48", Name: "Lozère"}, {Number: "49", Name: "Maine-et-Loire"}, {Number: "50", Name: "Manche"}, {Number: "51", Name: "Marne"}, {Number: "52", Name: "Haute-Marne"}, {Number: "53", Name: "Mayenne"}, {Number: "54", Name: "Meurthe-et-Moselle"}, {Number: "55", Name: "Meuse"}, {Number: "56", Name: "Morbihan"}, {Number: "57", Name: "Moselle"}, {Number: "58", Name: "Nièvre"}, {Number: "59", Name: "Nord"}, {Number: "60", Name: "Oise"}, {Number: "61", Name: "Orne"}, {Number: "62", Name: "Pas-de-Calais"}, {Number: "63", Name: "Puy-de-Dôme"}, {Number: "64", Name: "Pyrénées-Atlantiques"}, {Number: "65", Name: "Hautes-Pyrénées"}, {Number: "66", Name: "Pyrénées-Orientales"}, {Number: "67", Name: "Bas-Rhin"}, {Number: "68", Name: "Haut-Rhin"}, {Number: "69", Name: "Rhône"}, {Number: "70", Name: "Haute-Saône"}, {Number: "71", Name: "Saône-et-Loire"}, {Number: "72", Name: "Sarthe"}, {Number: "73", Name: "Savoie"}, {Number: "74", Name: "Haute-Savoie"}, {Number: "75", Name: "Paris"}, {Number: "76", Name: "Seine-Maritime"}, {Number: "77", Name: "Seine-et-Marne"}, {Number: "78", Name: "Yvelines"}, {Number: "79", Name: "Deux-Sèvres"}, {Number: "80", Name: "Somme"}, {Number: "81", Name: "Tarn"}, {Number: "82", Name: "Tarn-et-Garonne"}, {Number: "83", Name: "Var"}, {Number: "84", Name: "Vaucluse"}, {Number: "85", Name: "Vendée"}, {Number: "86", Name: "Vienne"}, {Number: "87", Name: "Haute-Vienne"}, {Number: "88", Name: "Vosges"}, {Number: "89", Name: "Yonne"}, {Number: "90", Name: "Territoire de Belfort"}, {Number: "91", Name: "Essonne"}, {Number: "92", Name: "Hauts-de-Seine"}, {Number: "93", Name: "Seine-Saint-Denis"}, {Number: "94", Name: "Val-de-Marne"}, {Number: "95", Name: "Val-d'Oise"}, {Number: "971", Name: "Guadeloupe"}, {Number: "972", Name: "Martinique"}, {Number: "973", Name: "Guyane"}, {Number: "974", Name: "La Réunion"}, {Number: "976", Name: "Mayotte"},}

var secondaryAddressFormats = []string{"Apt. ###", "Suite ###", "Étage ###", "Bât. ###", "Chambre ###"}

func SecondaryAddress() string {
	return faker.Numerify([]byte(faker.StringIn(secondaryAddressFormats)))
}

func CityPrefix() string {
	return faker.StringIn(citySuffix)
}

func StreetPrefix() string {
	return faker.StringIn(streetPrefix)
}

func BuildingNumber() string {
	return faker.Numerify([]byte(faker.StringIn(buildingNumber)))
}

func City() string {
	return faker.CompileRandValue(cityFormats, GetDataProviders())
}

func StreetName() string {
	return faker.CompileRandValue(streetNameFormats, GetDataProviders())
}

func StreetAddress() string {
	return faker.CompileRandValue(streetAddressFormats, GetDataProviders())
}

func Postcode() string {
	return strings.ToUpper(faker.Compute([]byte(faker.StringIn(postcode))))
}

func Address() string {
	return faker.CompileRandValue(addressFormats, GetDataProviders())
}

func Country() string {
	return faker.StringIn(country)
}

func Region() string {
	return faker.StringIn(regions)
}

func Department() string {
	return departments[faker.Intn(len(departments))].String()
}

func department() *Departement {
	return departments[faker.Intn(len(departments))]
}

func DepartmentName() string {
	return department().Name
}

func DepartmentNumber() string {
	return department().Number
}
