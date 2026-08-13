import { Layout } from "antd";

import Header from "../../components/AppHeader";
import InputPanel from "./InputPanel";
import ResultPanel from "./ResultPanel";
import HistoryTable from "./HistoryTable";

const { Content } = Layout;

export default function Parser() {
    return (
        <Layout className="app-layout">
            <Header />

            <Content className="app-content">
                {/* 输入区 + 解析结果区：响应式两栏，窄屏自动堆叠 */}
                <div className="main-grid">
                    <div className="left-panel">
                        <InputPanel />
                    </div>

                    <div className="right-panel">
                        <ResultPanel />
                    </div>
                </div>

                {/* 历史区域 */}
                <HistoryTable />

                <div className="app-footer">
                    © 2026 协议解析工具 | 让协议解析更简单高效
                </div>
            </Content>
        </Layout>
    );
}
