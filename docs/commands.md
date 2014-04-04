# Functions

## Roll(speed uint8, heading uint16)

This commands Sphero to roll along the provided vector. Both a speed
and a heading are required; the latter is considered relative to the last
calibrated direction.

#### Params

- **speed** - **uint8** - rotation speed
- **heading** - **uint16** - direction

#### API Command

**RollC**

## Stop

Stops the Sphero from rolling around.

#### API Command

**StopC**


## SetBackLED(level uint8)

This allows you to control the brightness of the back LED. The value does
not persist across power cycles.

#### Params

- **level** - **uint8** - brightness of back lED

#### API Command

**SetBackLEDC**

## SetRGB(r uint8, g uint8, b uint8)

Sets the sphero's LED color

#### Params

- **r** - **uint8** - Red
- **g** - **uint8** - Green
- **b** - **uintu** - Blue

#### API Command

**SetRGBC**

## GetRGB()

This retrieves the "user LED color" which is stored in the config block.

#### Returns 

- **[]uint8** - the sphero's LED color

#### API Command

**GetRGBC**

## SetStabalisation(on bool)

Sets whether the Sphero should have stabilization enabled

- **on** - **bool** - whether or not the sphero should have stabilization

#### API Command

**SetStabalisationC**

## SetHeading(heading uint16)

This allows the client to adjust the orientation of Sphero
by commanding a new reference heading in degrees, which ranges from 0 to 359.

#### Params

- **heading** - **uint16** - heading

#### API Command

**SetHeadingC**

