import {
    Card,
    Table,
    Button,
    Tag
} from "antd";


import {
    mockHistory
} from "../../mock/parser";



export default function HistoryTable() {



    const columns = [


        {
            title: "解析时间",

            dataIndex: "time",

            key: "time"

        },


        {
            title: "协议",

            dataIndex: "protocol",

            key: "protocol",

            render: (value: string) => (

                <Tag color="blue">

                    {value}

                </Tag>

            )

        },


        {
            title: "消息ID",

            dataIndex: "messageId",

            key: "messageId"

        },


        {
            title: "消息名称",

            dataIndex: "messageName",

            key: "messageName"

        },


        {
            title: "长度",

            dataIndex: "length",

            key: "length"

        },


        {
            title: "操作",

            key: "action",

            render: () => (

                <Button

                    type="link"

                >

                    查看

                </Button>

            )

        }



    ]



    return (

        <Card

            title="解析记录（历史）"

        >


            <Table


                columns={columns}


                dataSource={mockHistory}


                rowKey="id"


                pagination={false}


                size="small"


            />


        </Card>

    )

}