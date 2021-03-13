
wl = @website.website_locations.first

unless wl.extra_storage.positive?
  wl.extra_storage += 1
  wl.save
end

if @website.storage_areas.empty?
  @website.storage_areas = ["/data"]
  @website.save
end
