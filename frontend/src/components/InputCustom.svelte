<script>
    export let value;
    export let input_tipe = "text";
    export let input_disabled = false;
    export let input_maxlenght = 11;
    export let input_placeholder = "";

    const handleKeyboard_upppercase = (e) => {
        let str = e.target.value;
        let newstr = str.replace(" ","");
        return e.target.value = newstr.toUpperCase();
	}
</script>
{#if input_tipe == "text_uppercase_trim"}
    <input 
        bind:value
        on:keyup={handleKeyboard_upppercase}
        disabled={input_disabled}
        class="required form-control"
        maxlength="{input_maxlenght}"
        type="text"
        placeholder="{input_placeholder}"/>
{/if}