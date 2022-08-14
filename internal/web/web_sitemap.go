package web

import (
	"io"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
)

func (data Data) RobotsTxtHand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	// Generate robots.txt
	robotsTxt := "User-agent: *\nDisallow: /\n"

	if *data.RobotsDisallow == false {
		proto := netshare.GetProtocol(req.Header)
		host := netshare.GetHost(req)

		robotsTxt = "User-agent: *\nAllow: /\nSitemap: " + proto + "://" + host + "/sitemap.xml\n"
	}

	// Write response
	rw.Header().Set("Content-Type", "text/plain")
	_, err := io.WriteString(rw, robotsTxt)
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}

func (data Data) SitemapHand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	// Get protocol and host
	proto := netshare.GetProtocol(req.Header)
	host := netshare.GetHost(req)

	// Generate sitemap.xml
	sitemapXML := `<?xml version="1.0" encoding="UTF-8"?>`
	sitemapXML = sitemapXML + "\n" + `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n"
	sitemapXML = sitemapXML + "<url><loc>" + proto + "://" + host + "/" + "</loc></url>\n"
	sitemapXML = sitemapXML + "<url><loc>" + proto + "://" + host + "/about" + "</loc></url>\n"
	sitemapXML = sitemapXML + "<url><loc>" + proto + "://" + host + "/docs/apiv1" + "</loc></url>\n"
	sitemapXML = sitemapXML + "<url><loc>" + proto + "://" + host + "/docs/api_libs" + "</loc></url>\n"
	sitemapXML = sitemapXML + "</urlset>\n"

	// Write response
	rw.Header().Set("Content-Type", "text/xml")
	_, err := io.WriteString(rw, sitemapXML)
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}
