<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPage = [];
    let listCurr = [];
    let listBranch = [];
    let listDepartement = [];
    let search = "";
    let record = "";
    let record_message = "";
    let totalrecord = 0;
    let totalrecordall = 0;
    let totalpaging = 0;
    let perpage = 0;
    let page = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "CURR-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome(e) {
        listHome = [];
        listCurr = [];
        listBranch = [];
        listDepartement = [];
        const res = await fetch("/api/purchaserequest", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                purchaserequest_search: e,
                purchaserequest_status: e,
                purchaserequest_page : parseInt(page)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_listcurr = json.listcurr;
            let record_listbranch = json.listbranch;
            let record_listdepartement = json.listdepartement;
            record_message = json.message;
            perpage = json.perpage;
            totalrecordall = json.totalrecord;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["vendor_id"],
                            home_idcatevendor: record[i]["vendor_idcatevendor"],
                            home_name: record[i]["vendor_name"],
                            home_nmcatevendor: record[i]["vendor_nmcatevendor"],
                            home_pic: record[i]["vendor_pic"],
                            home_alamat: record[i]["vendor_alamat"],
                            home_email: record[i]["vendor_email"],
                            home_phone1: record[i]["vendor_phone1"],
                            home_phone2: record[i]["vendor_phone2"],
                            home_status: record[i]["vendor_status"],
                            home_status_css: record[i]["vendor_status_css"],
                            home_create: record[i]["vendor_create"],
                            home_update: record[i]["vendor_update"],
                        },
                    ];
                }
                listPage = [];
                for(var i=1;i<=totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
                        },
                    ];
                }
            }
            for (var i = 0; i < record_listcurr.length; i++) {
                listCurr = [
                    ...listCurr,
                    {
                        curr_id: record_listcurr[i]["curr_id"],
                    },
                ];
            }
            for (var i = 0; i < record_listbranch.length; i++) {
                listBranch = [
                    ...listBranch,
                    {
                        branch_id: record_listbranch[i]["branch_id"],
                        branch_name: record_listbranch[i]["branch_name"],
                    },
                ];
            }
            for (var i = 0; i < record_listdepartement.length; i++) {
                listDepartement = [
                    ...listDepartement,
                    {
                        departement_id: record_listdepartement[i]["departement_id"],
                        departement_name: record_listdepartement[i]["departement_name"],
                    },
                ];
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        listCurr = [];
        listBranch = [];
        listDepartement = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    const handleSearch = (e) => {
        search = e.detail.searchHome;
        initHome(search,"")
    };
    const handlePaging = (e) => {
        page = e.detail.page
        initHome("")
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handlePaging={handlePaging}
    on:handleSearch={handleSearch}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listPage}
    {listHome}
    {listCurr}
    {listBranch}
    {listDepartement}
    {totalrecord}/>
{/if}