@echo off
echo Hi, welcome to CMKIT.
echo ......
echo 服务已启动，非特殊情况不要关闭此窗口！
echo 可以按CTRL+C停止服务
echo =====================
cmkit.exe -db.passwd 12345678 -web.dir "../nginx/html/cmkit-web/assets/pictures/" >log.txt 2>&1
pause