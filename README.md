@@ -1,75 +1,125 @@
-# React + TypeScript + Vite
+# 协议解析工具 (Protocol Parser Web)
 
-This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.
+> 让协议解析更简单高效 —— 在线 HEX 报文解析器，支持 JT-808 / 2929 等协议的报文解析与可视化。
 
-Currently, two official plugins are available:
+![React](https://img.shields.io/badge/React-19-61DAFB?logo=react)
+![TypeScript](https://img.shields.io/badge/TypeScript-6-3178C6?logo=typescript)
+![Vite](https://img.shields.io/badge/Vite-8-646CFF?logo=vite)
+![Ant Design](https://img.shields.io/badge/Ant%20Design-6-1677FF?logo=antdesign)
+![Zustand](https://img.shields.io/badge/Zustand-5-000000)
 
-- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Oxc](https://oxc.rs)
-- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/)
+---
 
-## React Compiler
+## 项目简介
 
-The React Compiler is not enabled on this template because of its impact on dev & build performances. To add it, see [this documentation](https://react.dev/learn/react-compiler/installation).
+**协议解析工具** 是一个基于 Web 的报文解析应用，旨在帮助开发者和测试人员快速解析 HEX 格式的通信协议报文。  
+支持多协议选择、报文示例填充、字段级解析结果展示（表格 / JSON），并附带历史解析记录管理。
 
-## Expanding the ESLint configuration
+## 功能一览
 
-If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:
+- ✅ **多协议支持** — 内置 JT-808、2929 协议模板，可扩展
+- ✅ **HEX 报文输入** — 支持 HEX 字符串录入，实时显示字节数
+- ✅ **一键解析** — 点击「解析报文」即可得到字段级解析结果
+- ✅ **双视图结果展示** — 表格视图（逐字段展示偏移/长度/值） + JSON 视图（结构化数据）
+- ✅ **报文示例** — 内置常用报文示例（位置上传、心跳、注册等），一键填充
+- ✅ **解析记录** — 历史解析记录列表，支持搜索与筛选
+- ✅ **复制能力** — 字段值一键复制、解析结果一键复制
+- ✅ **响应式布局** — 窄屏自动堆叠，适配桌面与移动端
 
-```js
-export default defineConfig([
-  globalIgnores(['dist']),
-  {
-    files: ['**/*.{ts,tsx}'],
-    extends: [
-      // Other configs...
+## 技术栈
 
-      // Remove tseslint.configs.recommended and replace with this
-      tseslint.configs.recommendedTypeChecked,
-      // Alternatively, use this for stricter rules
-      tseslint.configs.strictTypeChecked,
-      // Optionally, add this for stylistic rules
-      tseslint.configs.stylisticTypeChecked,
+| 类别     | 技术                                                  |
+| -------- | ----------------------------------------------------- |
+| 框架     | React 19                                              |
+| 语言     | TypeScript 6                                          |
+| 构建工具 | Vite 8                                                |
+| UI 组件  | Ant Design 6                                          |
+| 状态管理 | Zustand 5                                             |
+| 包管理器 | npm / yarn                                            |
 
-      // Other configs...
-    ],
-    languageOptions: {
-      parserOptions: {
-        project: ['./tsconfig.node.json', './tsconfig.app.json'],
-        tsconfigRootDir: import.meta.dirname,
-      },
-      // other options...
-    },
-  },
-])
+## 快速开始
 
+```bash
+# 1. 克隆项目
+git clone https://github.com/your-username/protcol-parser-web.git
+cd protcol-parser-web
+
+# 2. 安装依赖
+npm install or yarn install  or pnpm install or yarn install --ignore-engines
+
+# 3. 启动开发服务器
+npm run dev or yarn dev or pnpm dev
+
+# 4. 打开浏览器访问
+open http://localhost:5173
 ```
 
-You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:
+### 构建生产版本
 
-```js
-// eslint.config.js
-import reactX from 'eslint-plugin-react-x'
-import reactDom from 'eslint-plugin-react-dom'
+```bash
+npm run build  or yarn build or pnpm build
+```
 
-export default defineConfig([
-  globalIgnores(['dist']),
-  {
-    files: ['**/*.{ts,tsx}'],
-    extends: [
-      // Other configs...
-      // Enable lint rules for React
-      reactX.configs['recommended-typescript'],
-      // Enable lint rules for React DOM
-      reactDom.configs.recommended,
-    ],
-    languageOptions: {
-      parserOptions: {
-        project: ['./tsconfig.node.json', './tsconfig.app.json'],
-        tsconfigRootDir: import.meta.dirname,
-      },
-      // other options...
-    },
-  },
-])
+构建产物输出至 `dist/` 目录，可直接部署至 Nginx / Vercel 等静态托管服务。
 
+## 使用指南
+
+### 基本流程
+
+1. **选择协议** — 在顶部下拉框中选择报文协议（如 JT-808）
+2. **输入 HEX 报文** — 在文本框中粘贴或输入 HEX 字符串，也可点击「报文示例」快速填充
+3. **点击解析** — 按下「解析报文」按钮，右侧面板即时显示解析结果
+4. **查看结果** — 可在「表格视图」中逐字段查看偏移量、长度、原始值与解析值；或切换至「JSON 视图」查看完整结构化数据
+5. **管理记录** — 历史解析记录自动保存至下方表格，支持按消息 ID / 名称 / 协议搜索
+
+### 示例报文
+
+应用内置以下常用 JT-808 示例报文，可直接点击填充：
+
+| 示例              | 消息 ID | 说明         |
+| ----------------- | ------- | ------------ |
+| 位置上传          | 0x0200  | GPS 位置数据 |
+| 报警信息          | 0x0801  | 报警事件上报 |
+| 终端心跳          | 0x0002  | 心跳保活     |
+| 终端注册          | 0x0100  | 终端注册请求 |
+| 终端注册应答      | 0x8100  | 注册应答     |
+| 位置批量上传      | 0x0704  | 批量位置上报 |
+
+## 项目结构
+
 ```
+src/
+├── App.tsx                          # 应用入口
+├── main.tsx                         # 渲染入口
+├── index.css                        # 全局样式
+├── components/
+│   ├── AppHeader/                   # 顶部导航栏
+│   │   └── index.tsx
+│   ├── ParseTable/                  # 字段级解析表格
+│   │   └── index.tsx
+│   └── RecordDrawer/                # 历史记录抽屉（可扩展）
+│       └── index.tsx
+├── pages/
+│   └── Parser/                      # 主解析页面
+│       ├── index.tsx                # 页面布局
+│       ├── InputPanel.tsx           # 输入面板
+│       ├── ResultPanel.tsx          # 结果面板
+│       └── HistoryTable.tsx         # 历史记录表格
+├── store/
+│   └── parser.ts                    # Zustand 状态管理
+└── mock/
+    └── parser.ts                    # Mock 数据（示例报文 / 历史记录）
+```
+
+## Roadmap
+
+- [ ] 真实协议解析引擎（当前使用 Mock 数据展示）
+- [ ] 自定义协议配置 / 导入
+- [ ] 国际化（i18n）支持
+- [ ] 报文对比功能
+- [ ] 暗色模式
+
+## 许可证
+
+[MIT](./LICENSE)
+
