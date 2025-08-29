# app.rb
require 'sinatra'
require "sinatra/reloader" if development?
require 'net/http'
require 'uri'


# /ruby now chains to the PHP service's /php endpoint via the Service ClusterIP.
get '/ruby' do
  host = ENV['ALL_LANGUAGE_PHP_SERVICE_HOST'] || ENV['PHP_SERVICE_HOST'] || 'all-language-php-lb'
  port = ENV['ALL_LANGUAGE_PHP_SERVICE_PORT'] || ENV['PHP_SERVICE_PORT'] || '80'
  target = "http://#{host}:#{port}/php"

  uri = URI.parse(target)
  begin
    res = Net::HTTP.get_response(uri)
    status res.code.to_i
    headers 'X-Upstream' => 'ruby->php', 'X-PHP-Target' => target
    body "Ruby /ruby -> PHP /php\nPHP said: #{res.body}\n"
  rescue => e
    status 502
    body "Ruby failed to reach PHP /php at #{target}: #{e.class} - #{e.message}\n"
  end
end

# Keep a direct probe route if you still want one:
get '/php' do
  host = ENV['ALL_LANGUAGE_PHP_SERVICE_HOST'] || ENV['PHP_SERVICE_HOST'] || 'all-language-php'
  port = ENV['ALL_LANGUAGE_PHP_SERVICE_PORT'] || ENV['PHP_SERVICE_PORT'] || '80'
  target = "http://#{host}:#{port}/"
  uri = URI.parse(target)
  begin
    res = Net::HTTP.get_response(uri)
    status res.code.to_i
    headers 'X-Upstream' => "ruby->php", 'X-PHP-Target' => target
    body "Ruby called PHP at #{target}\nPHP said: #{res.body}\n"
  rescue => e
    status 502
    body "Ruby failed to reach PHP at #{target}: #{e.class} - #{e.message}\n"
  end
end

