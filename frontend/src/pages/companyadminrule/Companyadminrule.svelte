<script>
    import Home from "./Home.svelte";
    import Entry from "./Entry.svelte";

    let listAdminrule = [];
    let listCompany = [];
    let sData = "";
    let record = "";
    let totalrecord = 0;
    let adminrule_idadmin = "";
    let adminrule_idcompany = "";
    let adminrule_nmrule = "";
    let adminrule_rule = "";
    export let table_header_font = "";
    export let table_body_font = "";
    let token = localStorage.getItem("token");
    let akses_page = false;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ADMINRULE-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initAdminrule();
        }
    }
    async function initAdminrule() {
        listAdminrule = [];
        listCompany = [];
        const res = await fetch("/api/companyadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({}),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_company = json.listcompany;
           
            let no = 0;
            if (record != null) {
                totalrecord = record.length;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listAdminrule = [
                        ...listAdminrule,
                        {
                            adminrule_no: no,
                            adminrule_idadmin: record[i]["companyadminrule_id"],
                            adminrule_idcompany: record[i]["companyadminrule_idcompany"],
                            adminrule_nmrule: record[i]["companyadminrule_nmrule"],
                            adminrule_rule: record[i]["companyadminrule_rule"],
                        },
                    ];
                }
            }
            for (var i = 0; i < record_company.length; i++) {
                listCompany = [
                    ...listCompany,
                    {
                        company_id: record_company[i]["company_id"],
                        company_name: record_company[i]["company_name"],
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
    const handleBackHalaman = () => {
        adminrule_idadmin = "";
        adminrule_rule = "";
        sData = "";
        listAdminrule = [];
        handleRefreshData("all");
    };
    const handleEditData = (e) => {
        adminrule_idadmin = e.detail.a;
        adminrule_idcompany = e.detail.b;
        adminrule_nmrule = e.detail.c;
        adminrule_rule = e.detail.d;
        console.log(e)
        sData = "Edit";
    };
    const handleRefreshData = (e) => {
        listAdminrule = [];
        listCompany = [];
        totalrecord = 0;
        setTimeout(function () {
            initAdminrule();
        }, 500);
    };
    initapp();
</script>

{#if akses_page == true}
    {#if sData == "" && adminrule_idadmin == ""}
        <Home
            on:handleEditData={handleEditData}
            on:handleRefreshData={handleRefreshData}
            {token}
            {table_header_font}
            {table_body_font}
            {listAdminrule}
            {listCompany}
            {totalrecord}
        />
    {:else}
        <Entry
            on:handleBackHalaman={handleBackHalaman}
            {sData}
            {token}
            {table_header_font}
            {table_body_font}
            {adminrule_idadmin}
            {adminrule_idcompany}
            {adminrule_nmrule}
            {adminrule_rule}
        />
    {/if}
{/if}
