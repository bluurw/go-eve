package support

var allPatterns = []map[string]string{
	JSFrameworksPatterns,
	CSSUIPatterns,
	WebTechnologyPatterns,
	MarketingAnalyticsPatterns,
	WPPluginsPatterns,
	WAFPatterns,
	CMSPatterns,
}


var JSFrameworksPatterns = map[string]string{
	"jQuery":          `src=["'][^"']*jquery[^"']*["']`,
	"React":           `React(?:DOM)?\.`,
	"Vue.js":          `vue(?:\.js)?["']?`,
	"Angular":         `angular(?:\.module|\.)?`,
	"Ember.js":        `Ember(?:\.|["'])`,
	"Backbone.js":     `Backbone(?:\.|["'])`,
	"Preact":          `Preact(?:\.|["'])`,
}


var CSSUIPatterns = map[string]string{
	"Bootstrap":       `(href|src)=["'][^"']*bootstrap[^"']*["']`,
	"Tailwind CSS":    `href=["'][^"']*tailwind(?:\.min)?\.css["']`,
	"Materialize":     `materialize(?:\.min)?\.css`,
	"Bulma":           `href=["'][^"']*bulma(?:\.min)?\.css["']`,
	"FontAwesome":     `href=["'][^"']*font[-]?awesome[^"']*["']`,
}


var WebTechnologyPatterns = map[string]string{
	"Next.js":         `__NEXT_DATA__`,
	"Nuxt.js":         `__NUXT__`,
}


var MarketingAnalyticsPatterns = map[string]string{
	"Google Tag Manager": `(?i)(https://www.googletagmanager.com/gtm\.js\?id=GTM-[A-Z0-9]+|GTM-[A-Z0-9]+|dataLayer\s*=)`,
	"Google Analytics":   `www\.google-analytics\.com/analytics\.js`,
	"Hotjar":             `static\.hotjar\.com`,
	"Segment":            `cdn\.segment\.com`,
}


var WPPluginsPatterns = map[string]string{
	"LazyLoad":       `(?i)(loading=["']lazy["']|data-src=|data-lazy=|lazyload(?:\.min)?\.js|class=["'][^"']*lazy[^"']*["'])`,
}


var WAFPatterns = map[string]string{
	"Cloudflare":     `cdn-cgi/`,
}


var CMSPatterns = map[string]string{
	"WordPress":          `(?i)(wp-(content|includes)|<meta name=["']generator["'] content=["']WordPress[^"']*["'])`,
	"Joomla":             `(?i)(\/media\/system\/js|<meta name=["']generator["'] content=["']Joomla[^"']*["'])`,
	"Drupal":             `(?i)(\/sites\/(default|all)\/|<meta name=["']Generator["'] content=["']Drupal[^"']*["'])`,
	"Concrete CMS":       `(?i)(concrete5|concrete\.js|concrete5\/themes)`,
	"SilverStripe":       `(?i)(silverstripe|framework\/core)`,
	"TYPO3":              `(?i)(typo3\/|t3lib\/|<meta name=["']generator["'] content=["']TYPO3[^"']*["'])`,
	"Plone":              `(?i)(Plone\/|<meta name=["']generator["'] content=["']Plone[^"']*["'])`,
	"Umbraco":            `(?i)(Umbraco\/|<meta name=["']generator["'] content=["']Umbraco[^"']*["'])`,
	"Neos CMS":           `(?i)(Neos\.|NeosCMS|fusion\/objects|typo3fluid)`,
	"Bolt CMS":           `(?i)(bolt\/theme|bolt\.js|Powered by Bolt)`,
	"Textpattern":        `(?i)(textpattern|<meta name=["']generator["'] content=["']Textpattern[^"']*["'])`,
	"SPIP":               `(?i)(ecrire\/|spip\.php|<meta name=["']generator["'] content=["']SPIP[^"']*["'])`,
	"Fork CMS":           `(?i)(fork-cms|<meta name=["']generator["'] content=["']Fork CMS[^"']*["'])`,
	"Redaxo":             `(?i)(redaxo|rex_article|<meta name=["']generator["'] content=["']REDAXO[^"']*["'])`,
	"WonderCMS":          `(?i)(wondercms|<meta name=["']generator["'] content=["']WonderCMS[^"']*["'])`,
	"Publify":            `(?i)(Publify|typo\/stylesheets)`,
	"Omeka":              `(?i)(omeka-(classic|s)|<meta name=["']generator["'] content=["']Omeka[^"']*["'])`,
	"Exponent CMS":       `(?i)(exponentcms|exponent_bootstrap)`,
	"Lektor":             `(?i)(lektor|Powered by Lektor)`,
	"Pelican":            `(?i)(pelican\.css|pelican\.js|Powered by Pelican)`,
	"Mezzanine":          `(?i)(mezzanine|mezzanine\.js)`,
	"Kotti":              `(?i)(kotti|kotti\.static)`,
	"FeinCMS":            `(?i)(feincms|Powered by FeinCMS)`,
	"Quokka CMS":         `(?i)(quokka|Powered by Quokka)`,
	"Wagtail":            `(?i)(wagtail|wagtailadmin|Powered by Wagtail)`,
	"Superdesk":          `(?i)(superdesk|contentapi|superdesk\.js)`,
	"Strapi":             `(?i)(strapi\/admin|<title>Strapi Admin|strapi\.io)`,
	"Payload CMS":        `(?i)(payloadcms|\/admin\/payload)`,
	"Squidex":            `(?i)(squidex|\/api\/content\/)`,
	"Roadiz":             `(?i)(roadiz|rz-admin)`,
	"WinterCMS":          `(?i)(wintercms|winter\.js|cms\/themes)`,
	"Rapido":             `(?i)(rapido-cms|Powered by Rapido)`,
	"Typemill":           `(?i)(typemill|typemill\.js|\/tm\/assets)`,
	"Sveltia CMS":        `(?i)(sveltia|Powered by Sveltia)`,
	"Manifest CMS":       `(?i)(manifestcms|manifest\.yml)`,
	"OpenCms":            `(?i)(opencms\/|Powered by OpenCms)`,
	"Tiki Wiki CMS":      `(?i)(tikiwiki|tiki-(index|page|login))`,
	"ImpressPages":       `(?i)(impresspages|ipContent)`,
	"ProcessWire":        `(?i)(processwire|\/site\/assets\/|\/wire\/)`,
	"MODX":               `(?i)(modx|assets\/components|<meta name=["']generator["'] content=["']MODX[^"']*["'])`,
	"ExpressionEngine":   `(?i)(expressionengine|EE\.|\/system\/ee)`,
	"Kentico CMS":        `(?i)(kentico|CMSModules|CMSPages)`,
	"Django CMS":         `(?i)(django\.cms|{% load cms_tags %}|django\.js)`,
	"Grav":               `(?i)(grav|user\/pages|\/system\/assets)`,
	"CMSimple":           `(?i)(cmsimple|cmsimple\.css)`,
	"eZ Publish":         `(?i)(ezpublish|ezjscore|eZ\.js)`,
	"UMI.CMS":            `(?i)(umi\.cms|UMI\.js|Powered by UMI)`,
	"Statamic":           `(?i)(statamic|\/cp\/|\/site\/content)`,
	"Koken":              `(?i)(koken|koken\.js|\/storage\/themes)`,
	"Movable Type":       `(?i)(movabletype|mt\.js|mt-static)`,
	"Simplébo":           `(?i)(simplebo|Powered by Simplébo)`,
	"XpressEngine":       `(?i)(xpressengine|xe\.js|modules\/admin)`,
	"Dynamicweb":         `(?i)(dynamicweb|dwcontent|dwapi)`,
	"DataLife Engine":    `(?i)(datalifeengine|engine\/classes)`,
	"DNN":                `(?i)(dnn|dotnetnuke|\/DesktopModules\/)`,
	"Webflow":            `(?i)(webflow|webflow\.js|Powered by Webflow)`,
	"HubSpot CMS Hub":    `(?i)(hs-script-loader|hs-analytics|Powered by HubSpot)`,
	"Cargo":              `(?i)(cargo\.site|cargo\.js|Powered by Cargo)`,
	"Strato Website":     `(?i)(strato-editor|Powered by Strato)`,
	"Sitecore":           `(?i)(sitecore|sc_default\.js)`,
	"WebNode":            `(?i)(webnode|Powered by Webnode)`,
	"Contao":             `(?i)(contao|TL_FILES|<meta name=["']generator["'] content=["']Contao[^"']*["'])`,
	"Tilda":              `(?i)(tilda\.css|tilda-block|Powered by Tilda)`,
	"Duda":               `(?i)(duda|dudaone|cdn\.duda\.co)`,
}
