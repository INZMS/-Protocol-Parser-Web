import { useMemo, useState } from "react";
import { Card, Table, Button, Tag, Input, Space, message } from "antd";
import { SearchOutlined, ReloadOutlined } from "@ant-design/icons";
import type { ColumnsType } from "antd/es/table";

import RecordDrawer from "../../components/RecordDrawer";

import { mockHistory, mockRecordDetails } from "../../mock/parser";

export default function HistoryTable() {
    const [drawerOpen, setDrawerOpen] = useState(false);

    const [current, setCurrent] = useState<any>(null);
    const [keyword, setKeyword] = useState("");
    //const [historyList, setHistoryList] = useState(mockHistory);
    const [page, setPage] = useState(1);
    const filtered = useMemo(() => {
        const kw = keyword.trim().toLowerCase();
        if (!kw) { return mockHistory; }
        return mockHistory.filter(
            (item) =>
                item.messageId.toLowerCase().includes(kw) ||
                item.messageName.toLowerCase().includes(kw) ||
                item.protocol.toLowerCase().includes(kw)
        );
    }, [keyword]);

    const columns: ColumnsType<any> = [
        { title: "解析时间", dataIndex: "time", key: "time" },
        {
            title: "协议",
            dataIndex: "protocol",
            key: "protocol",
            render: (value: string) => <Tag color="blue">{value}</Tag>
        },
        { title: "消息ID", dataIndex: "messageId", key: "messageId" },
        { title: "消息名称", dataIndex: "messageName", key: "messageName" },
        {
            title: "报文长度",
            dataIndex: "length",
            key: "length",
            render: (value: number) => `${value} Bytes`
        },
        {
            title: "操作",
            key: "action",
            width: 160,
            render: (_: any, record: any) => (
                <Space size={4}>
                    <Button type="link" size="small" style={{ padding: 0 }} onClick={() => {

                        setCurrent({
                            ...record,
                            ...(mockRecordDetails[record.id] || {})
                        });

                        setDrawerOpen(true);

                    }}>查看</Button>
                    <Button type="link" size="small" style={{ padding: 0 }}
                        onClick={() => {


                            navigator.clipboard.writeText(

                                JSON.stringify(

                                    record,

                                    null,

                                    2

                                )

                            );


                            message.success("记录复制成功");


                        }}>复制</Button>
                    <Button type="link" size="small" danger style={{ padding: 0 }}>删除</Button>
                </Space>
            )
        }
    ];


    return (
        <>
            <Card
                title="解析记录（历史）"
                size="small"
                extra={
                    <Space>
                        <Input
                            allowClear
                            size="small"
                            placeholder="搜索消息ID/名称/协议"
                            prefix={<SearchOutlined />}
                            value={keyword}
                            onChange={(e) => { setKeyword(e.target.value); setPage(1); }}
                            style={{ width: 220 }}
                        />
                        <Button
                            size="small"
                            type="primary"
                            icon={<ReloadOutlined />}
                            onClick={() => {

                                setKeyword("");
                                setPage(1);

                                message.success("已刷新");

                            }}
                        >
                            刷新
                        </Button>
                    </Space>
                }
            >
                <Table
                    columns={columns}
                    dataSource={filtered}
                    rowKey="id"
                    size="small"
                    scroll={{ x: "max-content" }}
                    pagination={{
                        current: page,

                        pageSize: 5,
                        size: "small",
                        showTotal: (total) => `共 ${total} 条`,
                        onChange: (current) => setPage(current)
                    }}
                />
            </Card>
            <RecordDrawer

                open={drawerOpen}

                data={current}

                onClose={() => {

                    setDrawerOpen(false)

                }}

            />


        </>
    );
}
