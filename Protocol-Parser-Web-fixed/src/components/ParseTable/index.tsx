import { Table, Button, Tooltip, Typography } from "antd";
import { CopyOutlined } from "@ant-design/icons";

import { useParserStore } from "../../store/parser";

const MAX_VISIBLE_CHARACTERS = 20;

function EllipsisCell({ value, mono = false }: { value: unknown; mono?: boolean }) {
    const text = String(value ?? "");
    const characters = Array.from(text);
    const isLong = characters.length > MAX_VISIBLE_CHARACTERS;
    const displayText = isLong
        ? `${characters.slice(0, MAX_VISIBLE_CHARACTERS).join("")}…`
        : text;

    return (
        <Tooltip title={isLong ? text : undefined} placement="topLeft">
            <Typography.Text className={mono ? "mono" : undefined}>
                {displayText}
            </Typography.Text>
        </Tooltip>
    );
}

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
        { title: "序号", dataIndex: "index", key: "index", width: 50 },
        { title: "字段名称", dataIndex: "name", key: "name", width: 110 },
        { title: "起始位置", dataIndex: "offset", key: "offset", width: 72 },
        { title: "长度", dataIndex: "length", key: "length", width: 58 },
        {
            title: "原始值(HEX)",
            dataIndex: "raw",
            key: "raw",
            width: "21%",
            render: (value: string) => <EllipsisCell value={value} mono />
        },
        {
            title: "解析值",
            dataIndex: "value",
            key: "value",
            width: "21%",
            render: (value: string) => <EllipsisCell value={value} />
        },
        {
            title: "说明",
            dataIndex: "description",
            key: "description",
            render: (value: string) => <span style={{ overflowWrap: "anywhere" }}>{value}</span>
        },
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
            tableLayout="fixed"
        />
    );
}
