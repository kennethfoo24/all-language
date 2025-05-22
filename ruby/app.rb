# app.rb
require 'sinatra'
require "sinatra/reloader" if development?

# Respond "Hello world!" on /
get '/ruby' do
  "Hello World from Ruby!"
end

