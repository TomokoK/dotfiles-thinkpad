package main

import(
  "image/png"
  "image/draw"
  "image"
  "io"
  "os"
  "fmt"
  "encoding/base64"
//  "encoding/binary"
)

func main() {
  tileSize := 16;
  inputFile,_ := os.Open("file.png")
  defer inputFile.Close()
  reader := io.Reader(inputFile)
  img,_ := png.Decode(reader)
  rgba := image.NewRGBA(img.Bounds())
  draw.Draw(rgba, rgba.Bounds(),  img, rgba.Bounds().Min, draw.Over)
  width, height := rgba.Bounds().Max.X, rgba.Bounds().Max.Y
  mapwidth := width / tileSize
  mapheight := height / tileSize
  var tileset []image.Image
  var tilemap []uint32
  
  for y := 0; y + tileSize <= height; y += tileSize {
    for x := 0; x + tileSize <= width; x += tileSize {
      tile := rgba.SubImage(image.Rect(x, y, x+tileSize, y+tileSize))
      matchIndex := matchAny(tileset, tile)
      if  matchIndex < 0 {
        tilemap = append(tilemap, uint32(len(tileset)))
        tileset = append(tileset, tile)
        fmt.Print(".")
      } else {
        tilemap = append(tilemap, uint32(matchIndex))
      }
    }
  }


  fmt.Println("")
  fmt.Printf("%d tiles found\n", len(tileset))
  fmt.Println("Writing TMX file...")
  writeTMX(tileset, tilemap, mapwidth, mapheight)
}

func matchAny(tileset []image.Image, img image.Image) int {
  for i := range tileset {
    if match(tileset[i], img) {
      return i
    }
  }
  return -1
}

func match(img1, img2 image.Image) bool {
  for x := 0; x < 16; x++ {
    for y := 0; y < 16; y++ {
      r1, g1, b1, a1 := img1.At(img1.Bounds().Min.X + x, img1.Bounds().Min.Y + y).RGBA()
      r2, g2, b2, a2 := img2.At(img2.Bounds().Min.X + x, img2.Bounds().Min.Y + y).RGBA()
      if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
        return false 
      }
    }
  }
  return true 
}

func writeTMX(tileset []image.Image, tilemap []uint32, mapwidth, mapheight int) error {

  tilesize := 16 // temporary
  outputFile,err := os.Create("test.tmx")
  if err != nil {
    return err
  }
  defer outputFile.Close()
  writer := io.Writer(outputFile)
  var encoder io.WriteCloser
  writer.Write(([]byte)(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
  <map version="1.0" orientation="orthogonal" width="%d" height="%d" tilewidth="%d" tileheight="%d">
   <tileset firstgid="1" name="tileset" tilewidth="%d" tileheight="%d">`,
  mapwidth, mapheight, tilesize, tilesize, tilesize, tilesize)))

  for index, tile := range tileset { 
    writer.Write(([]byte)("\n"))
    writer.Write(([]byte)(fmt.Sprintf(`<tile id="%d">
    <image width="%d" height="%d" format="png">
    <data encoding="base64">`, index, tilesize, tilesize)))
    encoder = base64.NewEncoder(base64.StdEncoding, writer)
    png.Encode(encoder, tile)
    encoder.Close()
    writer.Write(([]byte)(fmt.Sprintf(`</data></image></tile>`)))
  }
  writer.Write(([]byte)(fmt.Sprintf(`
  </tileset>
   <layer name="Tile Layer 1" width="%d" height="%d">
    <data encoding="csv">`, mapwidth, mapheight)))
  for i, index := range tilemap {
    if i > 0 {
      writer.Write(([]byte)(","))
    }
    if i % mapwidth == 0 {
      writer.Write(([]byte)("\n"))
    }
    writer.Write(([]byte)(fmt.Sprintf("%d",index+1)))

//    binary.Write(encoder, binary.LittleEndian, index)
  }
  writer.Write(([]byte)(`
    </data>
   </layer>
  </map>`))

  return nil
}
