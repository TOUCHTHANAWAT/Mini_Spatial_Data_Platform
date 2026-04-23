# Mini_Spatial_Data_Platform

โปรเจกต์นี้เป็น Web Application สำหรับจัดการและแสดงข้อมูลเชิงพื้นที่บนแผนที่ โดยใช้
- Frontend: React + Vite + Ant Design + MapLibre GL JS 
- Backend: Go(Echo)
- Database: MongoDB

# Backend Setup (Go + Echo + MongoDB)
## 1. Initialize Project
go mod init github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform

## 2. Install Dependencies
go get github.com/labstack/echo/v4
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv

## 3. Clean Dependencies
go mod tidy

## 4. Run Backend Server
go run main.go

Backend จะรันที่:
http://localhost:8080

# Frontend Setup (Vite + React + MapLibre + Ant Design)
## 1. Create Project
npm create vite@5 frontend
cd frontend

## 2. Install Dependencies
npm install
npm install maplibre-gl
npm install antd
npm install @ant-design/icons@6.x --save
npm install @mapbox/mapbox-gl-draw
npm install maplibre-gl-draw

## 3. Run Frontend
npm run dev

Frontend จะรันที่:
http://localhost:5173

# Run Full Project
## Backend
### go run main.go

## Frontend
### npm run dev

# Environment Variables

โปรเจกต์นี้ต้องใช้ไฟล์ `.env` ในแต่ละ service

### Backend (`backend/.env`)
.env
MONGO_URI=mongodb://localhost:27017
DB_NAME=spatial_db
PORT=8080

## สิ่งที่ทำได้
Backend:
- พัฒนา RESTful API สำหรับจัดการข้อมูลสถานที่ ดู (List/Get) เพิ่ม (Create) แก้ไข (Edit) และลบ (Delete)
- ข้อมูลแต่ละรายการต้องอยู่ในรูปแบบ GeoJSON Feature
- ออกแบบ API Path และ HTTP Method ให้เหมาะสมตามหลัก RESTful และ Response ทั้งหมดต้องอยู่ในรูปแบบ JSON

Frontend:
- เรียกข้อมูลจาก Backend API
- แสดงรายการสถานที่ในรูปแบบ ตาราง (Table)
- แสดงสถานที่บน แผนที่ Interactive ด้วย MapLibre
- สามารถ เพิ่ม แก้ไข และ ลบ ข้อมูลสถานที่ผ่านหน้าแอปพลิเคชันได้
- ค้นหา สถานที่ตามชื่อ และ กรองตาม category

## Bonus ที่ทำเพิ่ม
### Draw Area Search
#### จุดประสงค์
เพื่อให้ผู้ใช้สามารถวาดพื้นที่บนแผนที่เพื่อค้นหาข้อมูลสถานที่ภายในขอบเขตที่วาด

#### วิธีใช้งาน
- ผู้ใช้กดเครื่องมือวาด polygon tool ที่มุมบนขวาแผนที่
- กดจิ้มที่แผนที่เพื่อกำหนดขอบเขตที่ต้องการ จากนั้นกด Enter ระบบจะแสดงหมุดและรายชื่อสถานที่ที่อยู่ในขอบเขตนั้น
- เมื่อต้องการลบขอบเขตให้กดที่ขอบเขตนั้น และกด Delect ที่มุมบขวา แล้วสถานที่ทั้งหมดจะกลับมาเหมือนเดิม
