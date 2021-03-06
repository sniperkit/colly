package whois

//TODO: Update from list https://github.com/undiabler/golang-whois/blob/master/servers.go
var nics = map[string]string{
	".com":  "whois.verisign-grs.com",
	".net":  "whois.verisign-grs.com",
	".org":  "whois.pir.org",
	".info": "whois.afilias.info",
	".ac":   "whois.nic.ac",
	".ae":   "whois.aeda.net.ae",
	".aero": "whois.aero",
	".af":   "whois.nic.af",
	".ag":   "whois.nic.ag",
	".ai":   "whois.ai",
	".am":   "whois.amnic.net",
	".arpa": "whois.iana.org",
	".as":   "whois.nic.as",
	".asia": "whois.nic.asia",
	".at":   "whois.nic.at",
	".au":   "whois.audns.net.au",
	".be":   "whois.dns.be",
	".bg":   "whois.register.bg",
	".biz":  "whois.biz",
	".bj":   "whois.nic.bj",
	".bo":   "whois.nic.bo",
	".br":   "whois.registro.br",
	".ca":   "whois.cira.ca",
	".cat":  "whois.cat",
	".cc":   "ccwhois.verisign-grs.com",
	".ch":   "whois.nic.ch",
	".ci":   "whois.nic.ci",
	".cl":   "whois.nic.cl",
	".cn":   "whois.cnnic.cn",
	".coop": "whois.nic.coop",
	".cx":   "whois.nic.cx",
	".cz":   "whois.nic.cz",
	".de":   "whois.denic.de",
	".dk":   "whois.dk-hostmaster.dk",
	".dm":   "whois.nic.dm",
	".ec":   "whois.nic.ec",
	".edu": "whois.educause.edu 	",
	".ee": "whois.eenet.ee 	",
	".eu": "whois.eu 	",
	".fi": "whois.ficora.fi 	",
	".fr": "whois.nic.fr 	",
	".gd":     "whois.adamsnames.com",
	".gg":     "whois.gg",
	".gi":     "whois2.afilias-grs.net",
	".gl":     "whois.nic.gl",
	".gov":    "whois.dotgov.gov",
	".gs":     "whois.nic.gs",
	".gy":     "whois.registry.gy",
	".hk":     "whois.hkirc.hk",
	".hn":     "whois2.afilias-grs.net",
	".ht":     "whois.nic.ht",
	".ie":     "whois.domainregistry.ie",
	".il":     "whois.isoc.org.il",
	".im":     "whois.nic.im",
	".in":     "whois.inregistry.net",
	".int":    "whois.iana.org",
	".io":     "whois.nic.io",
	".ir":     "whois.nic.ir",
	".is":     "whois.isnic.is",
	".it":     "whois.nic.it",
	".je":     "whois.je",
	".jobs":   "jobswhois.verisign-grs.com",
	".jp":     "whois.jprs.jp",
	".ke":     "whois.kenic.or.ke",
	".ki":     "whois.nic.ki",
	".kp":     "whois.kcce.kp",
	".kr":     "whois.nic.or.kr",
	".kz":     "whois.nic.kz",
	".la":     "whois.nic.la",
	".li":     "whois.nic.li",
	".lt":     "whois.domreg.lt",
	".lu":     "whois.dns.lu",
	".lv":     "whois.nic.lv",
	".ly":     "whois.nic.ly",
	".ma":     "whois.iam.net.ma",
	".md":     "whois.nic.md",
	".me":     "whois.nic.me",
	".mg":     "whois.nic.mg",
	".mn":     "whois.nic.mn",
	".mobi":   "whois.dotmobiregistry.net",
	".mp":     "whois.nic.mp",
	".ms":     "whois.nic.ms",
	".mu":     "whois.nic.mu",
	".museum": "whois.museum",
	".mx":     "whois.mx",
	".my":     "whois.domainregistry.my",
	".na":     "whois.na-nic.com.na",
	".name":   "whois.nic.name",
	".ng":     "whois.nic.net.ng",
	".nl":     "whois.domain-registry.nl",
	".no":     "whois.norid.no",
	".nu":     "whois.nic.nu",
	".nz":     "whois.srs.net.nz",
	".pe":     "kero.yachay.pe",
	".pl":     "whois.dns.pl",
	".pm":     "whois.nic.pm",
	".pr":     "whois.nic.pr",
	".pro":    "whois.registrypro.pro",
	".pt":     "whois.dns.pt",
	".re":     "whois.nic.re",
	".ro":     "whois.rotld.ro",
	".ru":     "whois.ripn.net",
	".sa":     "whois.nic.net.sa",
	".sb":     "whois.nic.net.sb",
	".sc":     "whois2.afilias-grs.net",
	".se":     "whois.iis.se",
	".sg":     "whois.sgnic.sg",
	".sh":     "whois.nic.sh",
	".si":     "whois.arnes.si",
	".sk":     "whois.sk-nic.sk",
	".sm":     "whois.ripe.net",
	".sn":     "whois.nic.sn",
	".so":     "whois.nic.so",
	".st":     "whois.nic.st",
	".tc":     "whois.adamsnames.tc",
	".tel":    "whois.nic.tel",
	".tf":     "whois.nic.tf",
	".th":     "whois.thnic.co.th",
	".tk":     "whois.dot.tk",
	".tl":     "whois.nic.tl",
	".tm":     "whois.nic.tm",
	".to":     "whois.tonic.to",
	".tr":     "whois.nic.tr",
	".travel": "whois.nic.travel",
	".tv":     "tvwhois.verisign-grs.com",
	".tw":     "whois.twnic.net.tw",
	".ua":     "whois.ua",
	".ug":     "whois.co.ug",
	".uk":     "whois.nic.uk",
	".us":     "whois.nic.us",
	".uy":     "whois.nic.org.uy",
	".uz":     "whois.cctld.uz",
	".vc":     "whois2.afilias-grs.net",
	".ve":     "whois.nic.ve",
	".vg":     "whois.adamsnames.tc",
	".wf":     "whois.nic.wf",
	".ws":     "whois.website.ws",
	".yt":     "whois.nic.yt",
}
