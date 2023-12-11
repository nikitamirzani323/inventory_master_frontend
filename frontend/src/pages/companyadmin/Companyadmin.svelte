<script>
  import Home from "./Home.svelte";
  import dayjs from "dayjs";
  let listAdmin = [];
  let listAdminrule = [];
  let listCompany = [];
  let record = "";
  let totalrecord = 0;
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
        page: "ADMIN-VIEW",
      }),
    });
    const json = await res.json();
    if (json.status === 400) {
      logout();
    } else if (json.status == 403) {
      alert(json.message);
    } else {
      akses_page = true;
      initAdmin();
    }
  }
  async function initAdmin() {
    const res = await fetch("/api/companyadmin", {
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
      
      let recordlistrule = json.listrule;
      let recordlistcompany= json.listcompany;
      let no = 0;
      if (record != null) {
        totalrecord = record.length;
        for (var i = 0; i < record.length; i++) {
          no = no + 1;
          listAdmin = [
            ...listAdmin,
            {
              admin_no: no,
              admin_id: record[i]["companyadmin_id"],
              admin_idcompany: record[i]["companyadmin_idcompany"],
              admin_idrule: record[i]["companyadmin_idrule"],
              admin_tipe: record[i]["companyadmin_tipe"],
              admin_username: record[i]["companyadmin_username"],
              admin_rule: record[i]["companyadmin_nmrule"],
              admin_nama: record[i]["companyadmin_name"],
              admin_phone1: record[i]["companyadmin_phone1"],
              admin_phone2: record[i]["companyadmin_phone2"],
              admin_lastlogin: record[i]["companyadmin_lastlogin"],
              admin_lastipaddres: record[i]["companyadmin_ipaddress"],
              admin_status: record[i]["companyadmin_status"],
              admin_statuscss: record[i]["companyadmin_status_css"],
              admin_create: record[i]["companyadmin_create"],
              admin_update: record[i]["companyadmin_update"],
            },
          ];
        }
      }
      if (recordlistrule != null) {
        for (let i = 0; i < recordlistrule.length; i++) {
          listAdminrule = [...listAdminrule,
            {
              companyadminrule_id: recordlistrule[i]["companyadminrule_id"],
              companyadminrule_idcompany: recordlistrule[i]["companyadminrule_idcompany"],
              companyadminrule_nmrule: recordlistrule[i]["companyadminrule_nmrule"],
            },
          ];
        }
      }
      if (recordlistcompany != null) {
        for (let i = 0; i < recordlistcompany.length; i++) {
          listCompany = [...listCompany,
            {
              company_id: recordlistcompany[i]["company_id"],
              company_name: recordlistcompany[i]["company_name"],
            },
          ];
        }
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
    listAdmin = [];
    listAdminrule = [];
    listCompany = [];
    totalrecord = 0;
    setTimeout(function () {
      initAdmin();
    }, 500);
  };
  initapp();
</script>

{#if akses_page == true}
  <Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listAdmin}
    {listAdminrule}
    {listCompany}
    {totalrecord}
  />
{/if}
