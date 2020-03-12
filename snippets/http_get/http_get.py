from __future__ import print_function
import requests

# get request
resp = requests.get(url) #$0

# access/use/print headers
print(resp.headers)

# access/use/print content of response body
print(resp.content)
