import {
    Drawer,
    Descriptions,
    Divider,
    Table,
    Button,
    message,
    Space
} from "antd";


import {
    CopyOutlined
} from "@ant-design/icons";



interface Props {

    open: boolean;

    data: any;

    onClose: () => void;

}




export default function RecordDrawer({

    open,

    data,

    onClose

}: Props) {



    if (!data) {

        return (

            <Drawer
                title="记录详情"
                placement="right"
                width={520}
                open={open}
                onClose={onClose}
            />
        )
    }




    const columns = [


        {
            title: "字段名称",
            dataIndex: "name",
            key: "name"
        },


        {
            title: "原始值(Hex)",
            dataIndex: "hex",
            key: "hex"
        },


        {
            title: "解析值",
            dataIndex: "value",
            key: "value"
        },


        {
            title: "说明",
            dataIndex: "desc",
            key: "desc"
        },

        {

            title: "操作",

            key: "action",

            width: 80,


            render: (_: any, record: any) => (


                <Button

                    type="link"

                    size="small"

                    icon={<CopyOutlined />}


                    style={{

                        color: "#1677ff"

                    }}


                    onClick={() => {


                        navigator.clipboard.writeText(

                            String(record.value || "")

                        );


                        message.success(

                            `${record.name} 已复制`

                        );


                    }}

                >

                    复制

                </Button>


            )

        }


    ];




    return (

        <Drawer

            title="记录详情"

            placement="right"

            width="40%"

            open={open}

            onClose={onClose}


        >



            {/* 基础信息 */}

            <Descriptions

                column={1}

                size="small"

            >

                <Descriptions.Item label="协议">

                    {data.protocol}

                </Descriptions.Item>


                <Descriptions.Item label="消息ID">

                    {data.messageId}

                </Descriptions.Item>



                <Descriptions.Item label="消息名称">

                    {data.messageName}

                </Descriptions.Item>



                <Descriptions.Item label="解析时间">

                    {data.time}

                </Descriptions.Item>



                <Descriptions.Item label="报文长度">

                    {data.length} Bytes

                </Descriptions.Item>


            </Descriptions>





            <Divider />



            {/* HEX */}
            <div

                style={{

                    display: "flex",

                    justifyContent: "space-between",

                    alignItems: "center"

                }}
            >
                <h4>

                    原始报文（HEX）

                </h4>


                <Button

                    size="small"
                    type="primary"

                    icon={<CopyOutlined />}

                    onClick={() => {


                        navigator.clipboard.writeText(

                            data.rawHex || ""

                        );


                        message.success("HEX复制成功");


                    }}
                >


                    复制HEX

                </Button>

            </div>

            <pre

                style={{

                    background: "#f6f8fa",

                    padding: 12,

                    borderRadius: 6,

                    whiteSpace: "pre-wrap",

                    wordBreak: "break-all",

                    marginBottom: 0

                }}

            >

                {data.rawHex || "-"}

            </pre>







            {/* 表格 */}

            <h4>

                解析结果（表格）

            </h4>



            <Table

                size="small"

                bordered

                pagination={false}

                rowKey="name"

                columns={columns}

                dataSource={data.fields || []}


            />






            <Divider />




            {/* JSON */}


            <div

                style={{

                    display: "flex",

                    justifyContent: "space-between"

                }}

            >

                <h4>

                    解析结果（JSON）

                </h4>



                <Button

                    size="small"
                    type="primary"

                    icon={<CopyOutlined />}

                    onClick={() => {

                        navigator.clipboard.writeText(

                            JSON.stringify(

                                data.json,

                                null,

                                2

                            )

                        )
                        message.success("JSON复制成功")

                    }}

                >

                    复制JSON

                </Button>


            </div>



            <pre

                style={{

                    background: "#f6f8fa",

                    padding: 12,

                    borderRadius: 6,

                    maxHeight: 300,

                    overflow: "auto"

                }}

            >

                {

                    JSON.stringify(

                        data.json,

                        null,

                        2

                    )

                }


            </pre>



        </Drawer >


    )


}