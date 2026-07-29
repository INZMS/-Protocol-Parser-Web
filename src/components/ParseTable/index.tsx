import {
    Table,
    Button
} from "antd";


import {
    CopyOutlined
} from "@ant-design/icons";


import {
    useParserStore
} from "../../store/parser";




export default function ParseTable() {



    const {

        result

    } = useParserStore();






    const data = result.fields.map(


        (item: any, index: number) => (


            {

                key: index,

                ...item

            }


        )


    );









    const columns = [



        {

            title: "字段名称",

            dataIndex: "name",

            key: "name"

        },



        {

            title: "HEX值",

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



            render: (_: any, record: any) => (



                <Button


                    type="text"


                    icon={<CopyOutlined />}



                    onClick={() => {


                        navigator.clipboard.writeText(

                            record.value

                        )


                    }}



                />



            )


        }



    ];









    return (


        <div

            style={{

                height:"100%",

                overflow:"hidden"

            }}

        >


            <Table



                columns={columns}



                dataSource={data}



                pagination={false}



                size="small"



                bordered



                scroll={{


                    y:"calc(100vh - 430px)"


                }}



            />


        </div>



    )

}