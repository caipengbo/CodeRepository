
cmake_minimum_required(VERSION 3.5)

# Set the project name
project(KVStore)  # 会产生${PROJECT_NAME}变量

# Add an executable
add_executable(main main.cpp)  # 第一个参数是可执行文件的名字，第二个参数是源文件列表
