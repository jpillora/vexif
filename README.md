# vexif

View EXIF data with `vexif <file>`

[![CI](https://github.com/jpillora/vexif/workflows/CI/badge.svg)](https://github.com/jpillora/vexif/actions?workflow=CI)

### Install

**Binaries**

[![Releases](https://img.shields.io/github/release/jpillora/vexif.svg)](https://github.com/jpillora/vexif/releases)
[![Releases](https://img.shields.io/github/downloads/jpillora/vexif/total.svg)](https://github.com/jpillora/vexif/releases)

Find [the latest pre-compiled binaries here](https://github.com/jpillora/vexif/releases/latest)  or download and install it now with:

```
$ curl https://i.jpillora.com/vexif! | bash
```

**Source**

```sh
$ go get -v github.com/jpillora/vexif
```

### Usage

```sh
$ vexif -h

  Usage: vexif [options] [file] [file] ...

  file to view exif info

  Options:
  --help, -h  display help
```

### Example

```sh
$ vexif my-pic.jpg

>>> ./my-pic.jpg
  ApertureValue = "4845/1918"
  BrightnessValue = "456/5621"
  ColorSpace = 1
  ComponentsConfiguration = ""
  DateTime = "2013:05:17 20:08:37"
  DateTimeDigitized = "2013:05:17 20:08:37"
  DateTimeOriginal = "2013:05:17 20:08:37"
  ExifIFDPointer = 192
  ExifVersion = "0221"
  ExposureMode = 0
  ExposureProgram = 2
  ExposureTime = "1/30"
  FNumber = "12/5"
  Flash = 32
  FlashpixVersion = "0100"
  FocalLength = "77/20"
  FocalLengthIn35mmFilm = 35
  ISOSpeedRatings = 800
  Make = "Apple"
  MeteringMode = 5
  Model = "iPhone 4"
  Orientation = 6
  PixelXDimension = 640
  PixelYDimension = 480
  ResolutionUnit = 2
  SceneCaptureType = 0
  SensingMethod = 2
  ShutterSpeedValue = "7192/1463"
  Software = "6.1.3"
  ThumbJPEGInterchangeFormat = 656
  ThumbJPEGInterchangeFormatLength = 7091
  WhiteBalance = 0
  XResolution = "72/1"
  YCbCrPositioning = 1
  YResolution = "72/1"
>>> done
```
