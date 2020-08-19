# NodeJS Prometheus Client Bug
Bug where importing the `prom-client` NodeJS package crashes the process if 
there is no stdin.

# Table Of Contents
- [Overview](#overview)

# Overview
This is a minimum self contained example.

## The Bug
If the NodeJS `prom-client` library is imported when there is no standard in it
crashes on a `read` system call with the `ENOTCONN` error.

The stack trace looks like so:

```
+ exec node index.js
hello world
events.js:291
       throw er; // Unhandled 'error' event
       ^
 Error: read ENOTCONN
     at tryReadStart (net.js:574:20) 
     at ReadStream.Socket._read (net.js:585:5)
     at ReadStream.Readable.read (_stream_readable.js:475:10)
     at ReadStream.Socket.read (net.js:625:39)
     at ReadStream.Socket (net.js:377:12)
     at new ReadStream (tty.js:58:14)
     at process.getStdin [as stdin] (internal/bootstrap/switches/is_main_thread.js:151:15)
     at get (<anonymous>)
     at getOwn (internal/bootstrap/loaders.js:150:5)
     at NativeModule.syncExports (internal/bootstrap/loaders.js:258:31)
	 Emitted 'error' event on ReadStream instance at:
     at emitErrorNT (internal/streams/destroy.js:100:8)
     at emitErrorCloseNT (internal/streams/destroy.js:68:3)
     at processTicksAndRejections (internal/process/task_queues.js:80:21) {
	 errno: -107,
	 code: 'ENOTCONN',
	 syscall: 'read'
}
```

## Execute
To run this example:

```
npm install
node index.js
```

It should import the `prom-client` library, print `hello world`, and exit.

However if you run this as a non-interactive background service the error 
will occur. In my case I use [runit](http://smarden.org/runit/) as my
system initializer. Included is a `run` file, to install and run:

```
# mkdir /etc/sv/nodejs-prom-client
# cp ./run /etc/sv/nodejs-prom-client
# ln -s /etc/sv/nodejs-prom-client /var/services
# sv up nodejs-prom-client
```

You will see that the `nodejs-prom-client` service does not start. If you look
at logs (by configuring `/etc/sv/nodejs-prom-client/log/run`) you will see the
`ENOTCONN` error.
