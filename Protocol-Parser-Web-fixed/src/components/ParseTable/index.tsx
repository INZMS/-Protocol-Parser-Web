import { Button, Table, Tooltip, Typography, message } from "antd";
import { CopyOutlined } from "@ant-design/icons";

import { useParserStore } from "../../store/parser";

const MAX_VISIBLE_CHARACTERS = 20;

function CompactValue({ value, mono = false }: { value: unknown; mono?: boolean }) {
    const text = String(value ?? "");
    const characters = Array.from(text);
    const isLong = characters.length > MAX_VISIBLE_CHARACTERS;
    const visible = isLong
        ? `${characters.slice(0, MAX_VISIBLE_CHARACTERS).join("")}…`
        : text;

    return (
        <Tooltip title={isLong ? text : undefined} placement="topLeft">
            <Typography.Text className={mono ? "mono parse-cell-raw" : "parse-cell-value"}>
                {visible}
            </Typography.Text>
        </Tooltip>
    );
}

export default function ParseTable() {
    const { result } = useParserStore();

    if (!result) {
        return null;
    }

    const data = result.fields.map((item, index) => ({
        key: index,
        displayIndex: index + 1,
        ...item
    }));

    const columns = [
        {
            title: "字段",
            dataIndex: "name",
            key: "name",
            width: 150,
            render: (value: string, record: any) => (
                <div className="parse-field-name">
                    <span className="parse-field-index">{record.displayIndex}</span>
                    <span>{value}</span>
                </div>
            )
        },
        {
            title: "位置 / 长度",
            key: "position",
            width: 100,
            render: (_: unknown, record: any) => (
                <span className="parse-position">
                    {record.offset} <span>·</span> {record.length}B
                </span>
            )
        },
        {
            title: "原始值",
            dataIndex: "raw",
            key: "raw",
            width: 190,
            render: (value: string) => <CompactValue value={value} mono />
        },
        {
            title: "解析结果",
            dataIndex: "value",
            key: "value",
            width: 230,
            render: (value: string) => (
                <div className="parse-value-with-action">
                    <CompactValue value={value} />
                    <Tooltip title="复制完整解析值">
                        <Button
                            className="parse-copy-button"
                            type="text"
                            size="small"
                            icon={<CopyOutlined />}
                            aria-label="复制完整解析值"
                            onClick={() => {
                                navigator.clipboard.writeText(String(value ?? ""));
                                message.success("解析值已复制");
                            }}
                        />
                    </Tooltip>
                </div>
            )
        },
        {
            title: "说明",
            dataIndex: "description",
            key: "description",
            render: (value: string) => <span className="parse-description">{value}</span>
        }
    ];

    return (
        <Table
            className="parse-result-table"
            columns={columns}
            dataSource={data}
            pagination={false}
            size="small"
            tableLayout="fixed"
            rowClassName={(_, index) => index % 2 === 1 ? "parse-row-even" : ""}
        />
    );
}
