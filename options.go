package main

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// Configuration Options that can be set by Command Line Flag, or Config File
type Options struct {
	HttpAddress             string        `flag:"http-address" cfg:"http_address"`
	RedirectUrl             string        `flag:"redirect-url" cfg:"redirect_url"`
	ClientID                string        `flag:"client-id" cfg:"client_id"`
	ClientSecret            string        `flag:"client-secret" cfg:"client_secret"`
	PassBasicAuth           bool          `flag:"pass-basic-auth" cfg:"pass_basic_auth"`
	HtpasswdFile            string        `flag:"htpasswd-file" cfg:"htpasswd_file"`
	CookieSecret            string        `flag:"cookie-secret" cfg:"cookie_secret"`
	CookieDomain            string        `flag:"cookie-domain" cfg:"cookie_domain"`
	CookieExpire            time.Duration `flag:"cookie-expire" cfg:"cookie_expire"`
	CookieHttpsOnly         bool          `flag:"cookie-https-only" cfg:"cookie_https_only"`
	AuthenticatedEmailsFile string        `flag:"authenticated-emails-file" cfg:"authenticated_emails_file"`
	GoogleAppsDomains       []string      `flag:"google-apps-domain" cfg:"google_apps_domains"`
	Upstreams               []string      `flag:"upstream" cfg:"upstreams"`

	// internal values that are set after config validation
	redirectUrl *url.URL
	proxyUrls   []*url.URL
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Validate() error {
	if len(o.Upstreams) < 1 {
		return errors.New("missing -upstream")
	}
	if o.CookieSecret == "" {
		errors.New("missing -cookie-secret")
	}
	if o.ClientID == "" {
		return errors.New("missing -client-id")
	}
	if o.ClientSecret == "" {
		return errors.New("missing -client-secret")
	}

	redirectUrl, err := url.Parse(o.RedirectUrl)
	if err != nil {
		return fmt.Errorf("error parsing -redirect-url=%q %s", o.RedirectUrl, err)
	}
	o.redirectUrl = redirectUrl

	for _, u := range o.Upstreams {
		upstreamUrl, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("error parsing -upstream=%q %s", upstreamUrl, err)
		}
		if upstreamUrl.Path == "" {
			upstreamUrl.Path = "/"
		}
		o.proxyUrls = append(o.proxyUrls, upstreamUrl)
	}

	return nil
}