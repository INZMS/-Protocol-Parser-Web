import { Table, Button } from "antd";
import { CopyOutlined } from "@ant-design/icons";

import { useParserStore } from "../../store/parser";

export default function ParseTable() {
    const { result } = useParserStore();

    if (!result) {
        return null;
    }

    const data = result.fields.map((item: any, index: number) => ({
        key: index,
        index: index + 1,
        ...item
    }));

    const columns = [
        { title: "序号", dataIndex: "index", key: "index", width: 56 },
        { title: "字段名称", dataIndex: "name", key: "name" },
        { title: "起始位置", dataIndex: "offset", key: "offset", width: 84 },
        { title: "长度(字节)", dataIndex: "len", key: "len", width: 92 },
        {
            title: "原始值(HEX)",
            dataIndex: "hex",
            key: "hex",
            render: (value: string) => <span className="mono">{value}</span>
        },
        { title: "解析值", dataIndex: "value", key: "value" },
        { title: "说明", dataIndex: "desc", key: "desc" },
        {
            title: "操作",
            key: "action",
            width: 60,
            render: (_: any, record: any) => (
                <Button
                    type="text"
                    size="small"
                    icon={<CopyOutlined />}
                    onClick={() => navigator.clipboard.writeText(record.value)}
                />
            )
        }
    ];

    return (
        <Table
            columns={columns}
            dataSource={data}
            pagination={false}
            size="small"
            bordered
        />
    );
}
