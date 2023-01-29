<script lang="ts">
    import type {CVESummary} from "./types"

    export let cveSummary: CVESummary

    let color = "text-neutral-800 bg-neutral-400";
    switch (cveSummary.cvssSeverity) {
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
            <span class="text-xl font-bold">{cveSummary.id}</span>
        </div>
        <div class="text-right">
            <div>
                <span class="p-1 text-sm rounded-md {color}">
                    {(cveSummary.cvssScore > 0.0 ? cveSummary.cvssScore + " " : "")}{cveSummary.cvssSeverity}
                </span>
            </div>
            <div>
                <span class="text-sm">{cveSummary.published}</span>
            </div>
        </div>
    </div>
    <div class="text-justify">
        <span>{cveSummary.desription}</span>
    </div>
</div>
