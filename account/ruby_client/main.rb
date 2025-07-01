#!/usr/bin/env ruby
# frozen_string_literal: true

puts "[Ruby gRPC Client] Starting request..."

# XXX: In the real world a Ruby gem would be generated and required instead
lib_dir = File.expand_path(File.join('..', 'gen', 'ruby'))
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'account/v1/account_services_pb'

# Create a stub for the service
stub = Account::V1::AccountService::Stub.new('localhost:8080', :this_channel_is_insecure)

# Build the request
request = Account::V1::PostAccountRequest.new(name: "Ruby")

# Call the RPC
begin
  response = stub.post_account(request)
  puts "Received response: #{response.inspect}"
rescue GRPC::BadStatus => e
  puts "GRPC Error: #{e.message}"
end
