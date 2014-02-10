# Events

## update

This is a meta-event that will be triggered when any of the other events are. If
you're looking for specific information, you might wish to subscribe to one of
the specific events instead.

## collision

Assuming that collision detection has been enabled, gets emitted when the sphero
hits something, falls from a step, or otherwise detects a collision.

## messages(data)

Gets triggered when the Sphero has data to send to the user.

## notifications(data)

Gets triggered when the Sphero has a notification to send to the user.

## connect

Gets triggered when the connection to the Sphero over Bluetooth has been opened.

## start

Gets triggered when the Sphero is started and ready to be used.