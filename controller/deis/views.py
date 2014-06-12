import json

from django.http import HttpResponse

from deis import __version__

def info(req):
    data = {'version': __version__}
    return HttpResponse(json.dumps(data))
