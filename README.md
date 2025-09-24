# 时间表计时器 (Schedule Timer)

## 简介 (Introduction / Введение)

这是一个基于 Go 语言的命令行时间表计时器程序，用于帮助用户根据 YAML 配置文件中的时间表管理任务和休息时间。程序会加载一个包含任务时间段的 YAML 文件，显示每个任务的倒计时，并在任务之间计算并显示休息时间。

This is a command-line schedule timer program written in Go, designed to help users manage tasks and break times based on a YAML configuration file. The program loads a YAML file containing task schedules, displays a countdown for each task, and calculates and displays break times between tasks.

Это программа-таймер расписания на основе командной строки, написанная на языке Go, предназначенная для помощи пользователям в управлении задачами и перерывами на основе конфигурационного файла YAML. Программа загружает файл YAML с расписанием задач, отображает обратный отсчет для каждой задачи и рассчитывает и отображает время перерывов между задачами.

## 功能 (Features / Функции)

- 从 YAML 文件加载时间表配置。
- 为每个任务显示倒计时。
- 在任务之间计算并显示休息时间。
- 使用简单的命令行界面，屏幕会定期清空以保持清晰。
- 支持错误处理，例如无效的时间格式。

- Loads schedule configuration from a YAML file.
- Displays a countdown timer for each task.
- Calculates and displays break times between tasks.
- Uses a simple command-line interface, clearing the screen periodically for clarity.
- Supports error handling, such as invalid time formats.

- Загружает конфигурацию расписания из файла YAML.
- Отображает таймер обратного отсчета для каждой задачи.
- Рассчитывает и отображает время перерывов между задачами.
- Использует простой интерфейс командной строки, периодически очищая экран для ясности.
- Поддерживает обработку ошибок, например, неверные форматы времени.

## 安装 (Installation / Установка)

1. 确保已安装 Go（版本 1.16 或更高）。
2. 克隆此仓库：
   ```bash
   git clone <repository-url>
   ```
3. 进入项目目录：
   ```bash
   cd schedule-timer
   ```
4. 安装依赖：
   ```bash
   go mod tidy
   ```
5. 构建并运行程序：
   ```bash
   go build
   ./schedule-timer --load-from-file <path-to-yaml-file>
   ```

1. Ensure Go (version 1.16 or higher) is installed.
2. Clone this repository:
   ```bash
   git clone <repository-url>
   ```
3. Navigate to the project directory:
   ```bash
   cd schedule-timer
   ```
4. Install dependencies:
   ```bash
   go mod tidy
   ```
5. Build and run the program:
   ```bash
   go build
   ./schedule-timer --load-from-file <path-to-yaml-file>
   ```

1. Убедитесь, что установлен Go (версия 1.16 или выше).
2. Склонируйте этот репозиторий:
   ```bash
   git clone <repository-url>
   ```
3. Перейдите в директорию проекта:
   ```bash
   cd schedule-timer
   ```
4. Установите зависимости:
   ```bash
   go mod tidy
   ```
5. Соберите и запустите программу:
   ```bash
   go build
   ./schedule-timer --load-from-file <path-to-yaml-file>
   ```

## 使用方法 (Usage / Использование)

1. 创建一个 YAML 配置文件（例如 `schedule.yaml`），格式如下：
   ```yaml
   Name: My Daily Schedule
   Tasks:
     - id: 1
       content: Morning Meeting
       start: 09:00
       end: 10:00
     - id: 2
       content: Coding Session
       start: 10:15
       end: 12:00
   ```

2. 运行程序并指定 YAML 文件：
   ```bash
   ./schedule-timer --load-from-file schedule.yaml
   ```

3. 程序将显示每个任务的倒计时，并在任务之间显示休息时间（如果有）。

1. Create a YAML configuration file (e.g., `schedule.yaml`) with the following format:
   ```yaml
   Name: My Daily Schedule
   Tasks:
     - id: 1
       content: Morning Meeting
       start: 09:00
       end: 10:00
     - id: 2
       content: Coding Session
       start: 10:15
       end: 12:00
   ```

2. Run the program with the YAML file:
   ```bash
   ./schedule-timer --load-from-file schedule.yaml
   ```

3. The program will display a countdown for each task and show break times (if any) between tasks.

1. Создайте конфигурационный файл YAML (например, `schedule.yaml`) в следующем формате:
   ```yaml
   Name: My Daily Schedule
   Tasks:
     - id: 1
       content: Morning Meeting
       start: 09:00
       end: 10:00
     - id: 2
       content: Coding Session
       start: 10:15
       end: 12:00
   ```

2. Запустите программу с указанием файла YAML:
   ```bash
   ./schedule-timer --load-from-file schedule.yaml
   ```

3. Программа будет отображать обратный отсчет для каждой задачи и показывать время перерывов (если они есть) между задачами.

## 依赖 (Dependencies / Зависимости)

- [github.com/inancgumus/screen](https://github.com/inancgumus/screen)：用于清屏和光标控制。
- [github.com/spf13/viper](https://github.com/spf13/viper)：用于解析 YAML 配置文件。
- [github.com/urfave/cli/v3](https://github.com/urfave/cli)：用于处理命令行参数。

- [github.com/inancgumus/screen](https://github.com/inancgumus/screen): For screen clearing and cursor control.
- [github.com/spf13/viper](https://github.com/spf13/viper): For parsing YAML configuration files.
- [github.com/urfave/cli/v3](https://github.com/urfave/cli): For handling command-line arguments.

- [github.com/inancgumus/screen](https://github.com/inancgumus/screen): Для очистки экрана и управления курсором.
- [github.com/spf13/viper](https://github.com/spf13/viper): Для разбора конфигурационных файлов YAML.
- [github.com/urfave/cli/v3](https://github.com/urfave/cli): Для обработки аргументов командной строки.

## 示例输出 (Example Output / Пример вывода)

```
1:23:45 Left
```

完成后，如果有休息时间：
```
BE BACK SOOOOON:
0:15:00
```

After a task completes, if there is a break:
```
BE BACK SOOOOON:
0:15:00
```

После завершения задачи, если есть перерыв:
```
BE BACK SOOOOON:
0:15:00
```

## 注意事项 (Notes / Примечания)

- 时间格式必须为 `HH:MM`（例如，`09:00`）。
- 确保 YAML 文件格式正确，否则程序可能无法正确解析。
- 程序会清空屏幕以保持界面简洁。

- Time format must be `HH:MM` (e.g., `09:00`).
- Ensure the YAML file is correctly formatted, or the program may fail to parse it.
- The program clears the screen to maintain a clean interface.

- Формат времени должен быть `HH:MM` (например, `09:00`).
- Убедитесь, что файл YAML правильно отформатирован, иначе программа может не разобрать его.
- Программа очищает экран для поддержания чистого интерфейса.

## 许可证 (License / Лицензия)

本项目采用 MIT 许可证。

This project is licensed under the MIT License.

Этот проект распространяется под лицензией MIT.