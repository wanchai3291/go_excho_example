# go_echo_example
## Golang using echo framework Template

#### เขียนด้วย concept modular

### Files
 - cmd -> main app
 - package(pkg) -> external file ไฟล์ที่ใช้รวมกันจะอยู่ในนี้เช่น random number service
 - src -> internal file จะทำงานเป็น module แยกขาดกันไม่ควร import ข้าม module จะทำให้ manage ยาก
 - config -> จะเก็บจำพวก config ต่างๆ เช่น databas  cache
 - assets -> เก็บไฟล์ static เช่น รูป html
 - scripts -> เก็บไฟล์ deploy เช่น docker ks8 docker-compose
 - test -> ไฟล์ test
