### basic types

1. bool
2. string
3. int  
   1. int8, int16, int32 (alias rune - represents a Unicode code point), int64 
4. uint 
   1. uint8(alias byte) uint16 uint32 uint64 uintptr (represents a Unicode code point)
5. float32, float64
6. complex64, complex128 

### Composite Types

1. pointer - C pointer alike.
2. struct - C struct alike.
3. function (func) - functions are first-class types in Go.
4. Containers
   3. array - fixed-length container types.
   4. slice - dynamic-length and dynamic-capacity container types
   5. map - maps are associative arrays (or dictionaries). The standard Go compiler implements maps as hashtables.
6. channel - channels are used to synchronize data among goroutines (the green threads in Go).
7. interface - interfaces play a key role in reflection and polymorphism. 

### zero values (аналог null, но не он)
у каждого типа свои. строки, числа, булы - понятно,
а вот у указателей, интерфейсов и ... устанавливается nil.
Отсюда вытекает проблема с обработкой null-значений из sql, json. есть разного рода трюки, например https://iamdual.com/en/posts/handle-sql-null-golang/ 
