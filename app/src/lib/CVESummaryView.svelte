<script lang="ts">
    export let cveID: string
    export let cveDescription: string
    export let published: string
    export let cvssScore: number
    export let cvssSeverity: string

    let color = "text-neutral-800 bg-neutral-400";
    switch (cvssSeverity) {
        case "LOW":
            color = "text-neutral-700 bg-yellow-300";
            break;
        case "MEDIUM":
            color = "text-neutral-700 bg-orange-500";
            break;
        case "HIGH":
            color = "text-neutral-100 bg-red-700";
            break;
        case "CRITICAL":
            color = "text-neutral-100 bg-red-900";
            break;
    }
</script>

<div class="bg-neutral-700 p-2 rounded-lg">
    <div class="grid grid-cols-2 mb-2">
        <div>
            <span class="text-xl font-bold">{cveID}</span>
        </div>
        <div class="text-right">
            <div>
                <span class="p-1 text-sm rounded-md {color}">
                    {(cvssScore > 0.0 ? cvssScore.toFixed(1) + " " : "")}{(cvssSeverity != "" ? cvssSeverity : "NO CVSS")}
                </span>
            </div>
            <div>
                <span class="text-sm">{published}</span>
            </div>
        </div>
    </div>
    <div>
        <span>{cveDescription}</span>
    </div>
</div>
