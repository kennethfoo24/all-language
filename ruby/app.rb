# app.rb
require 'sinatra'
require 'net/http'
require 'uri'
require 'datadog'
require 'datadog/auto_instrument'

Datadog.configure do |c|
  c.service = 'all-language-ruby'
  c.env = 'dev'
  c.version = '1.0'
  c.tracing.instrument :sinatra
  c.tracing.instrument :http
end

PHP_URL = URI('http://all-language-php-lb:80/php')

def http_get(uri)
  Net::HTTP.start(uri.host, uri.port) { |http| http.get(uri.request_uri) }
rescue => e
  e # return the exception so we can format a 502
end

get '/ruby' do
  res = http_get(PHP_URL)

  if res.is_a?(Net::HTTPResponse)
    status res.code.to_i
    headers 'X-Upstream' => 'ruby->php', 'X-PHP-Target' => PHP_URL.to_s
    "Ruby /ruby -> PHP /php\nPHP said: #{res.body}\n"
  else
    status 502
    "Ruby failed to reach PHP /php at #{PHP_URL}: #{res.class} - #{res.message}\n"
  end
end
