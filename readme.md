Step create file style hexagonal architecture

1. สร้างไฟล์เริ่มต้น
   1.1 สร้างไฟล์ app/cmd/api/main.go
   1.1 สร้าง Fiber application
   1.2 สร้าง GET endpoint เพื่อทดสอบว่า Project Start สำเร็จ
   1.3 เริ่ม Server ด้วย app.Listen(:port)
2. สร้าง .env และใส่ข้อมูล
3. สร้าง Config
   3.1 สร้างไฟล์ app/internal/config/config.go เพื่อ map ค่าจาก .env เข้าตัวแปร
   3.2 ประกาศ Config struct ตามค่า .env
   3.3 สร้าง Function LoadConfig() เพื่อโหลดค่าจาก .env มาเก็บใน Config struct
   3.4 เลือกใช้ Library godotenv เพื่อโหลดค่าจาก .env
4. สร้าง Entity ของ News
   4.1 สร้างไฟล์ app/internal/core/domain/entities/news.go
   4.2 ประกาศ News struct สำหรับเก็บข้อมูลข่าว
   4.3 กำหนด Struct Tag เพื่อแปลง field เป็น json
5. สร้าง Handler เพื่อเอาไปผูกกับ endpoint
   5.1 สร้างไฟล์ app/internal/adapters/http/handlers/news_handler.go
   5.2 ประกาศ NewsHandler struct
   5.3 สร้าง Constructor สำหรับรับ NewsService
   5.4 สร้าง Function GetByID() เพื่อรับ Request เหมือนตอนเขียนใน main.go แบบไฟล์เดียว
   5.5 อ่าน ID จาก URL Parameter
   5.6 เตรียมเรียก NewsService เพื่อดึงข้อมูลข่าว
6. 6.1 สร้าง app/internal/core/services/news_service.go
   6.2 ประกาศ NewsService struct
   6.3 สร้าง Constructor สำหรับรับ NewsRepository
   6.4 สร้าง Function GetByID() ไว้ก่อน โดยยังไม่ต้องเรียก Repository จริง
