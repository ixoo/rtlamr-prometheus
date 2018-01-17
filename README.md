# rtlamr-prometheus

A prometheus exporter for rtlamr SCM messages.

### Building
You need go language installed, then get the source and build using:

	go get github.com/ixoo/rtlamr-prometheus

Then feed it with the output of [`github.com/bemasher/rtlamr`]:

'''
rtlamr -format=json | ./rtlamr-prometheus
'''

Then configure your prometheus instance to scrap metrics on hostname:8080
