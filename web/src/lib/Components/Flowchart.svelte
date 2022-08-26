<script>
    import { hierarchy, tree } from "d3-hierarchy";
    import { linkVertical } from "d3-shape";
    
    export let data;

    const width = "100%";
    const height = "94%";

    const margin = {
        top: 60,
        left: 0,
        right: 0,
        bottom: 0,
    };

    const createData = (data, d) => {
        data.spec.steps.map((step) => {
            d.data = {
                name: step.stage.name,
                status: step.stage.status,
            }
        })

        // each step contains some steps
        // this is in form of array 
        // to create a graph for that 
        // a -> b -> c would be the right approach beacuse b will only run after a gets completed
        // to create a chart for that d3 need the data in form
        // to append b to a, a should contain a field called children, which is an array,
        // to append c to b, b should contain a field called children, which is an array,

        // data = [{}, {}, {}]
        // d = {
        //     children: [
        //         {
        //             children: [
        //                 {
        //                     children: []       
        //                 }
        //             ]                    
        //         }
        //     ]
        // }

        for(let i = 0; i < data.spec.steps.length; i++) {
            // main steps
            let step = data.spec.steps[i];

            // using t as a temporary object, for each step
            let t = {
                data: {},
                children: []
            };

            for(let j = step.stage.steps.length - 1; j >= 0; j--) {
                // steps in a step
                let subStep = step.stage.steps[j];
            
                // temperary object which will replace the current data(d)
                let temp = {
                    data: {},
                    children: []
                };

                // for first time substep is the last element 
                // which means there will be no children of that component
                if (j !== step.stage.steps.length - 1) {
                    temp.children.push(t);                    
                }

                // add all the data of the substep to the data field of the object
                temp.data = subStep;
                // updating t for each substep 
                t = temp;
            }

            // adding the object for each step in the data, containing 
            d.children.push(t);
        }
    }

    let d = {
        data: {},
        children: [],
    };

    createData(data, d);
    const treeLayout = tree().size([500, 600]);
    const root = hierarchy(d);
    const links = treeLayout(root).links();
    const pathGen = linkVertical().x(d => d.x).y(d => d.y);
    const paths = links.map(link => pathGen(link));
</script>

<main class="w-full xl:w-1/2">
    <svg {width} {height} class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <g transform={`translate(${margin.left},${margin.top})`}>
            {#each paths as path}
                <path d={path} fill="none" stroke="#fff"></path>
            {/each}
            <g stroke-linejoin="round" stroke-width=3>
                {#each root.descendants() as item}
                    <g transform={`translate(${item.x},${item.y})`}>
                        <rect fill="#ddd" x="-5rem" y="-1rem" height="2rem" width="11rem" rx="1rem"></rect>
                        {#if item.data.data.status === 'TimedOut' || item.data.data.status === 'Cancelled'}
                            <svg fill="red" x="-9.5rem" y="-0.62rem" height="1.25rem" width="11rem" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--! Font Awesome Pro 6.1.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc. --><path d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM175 208.1L222.1 255.1L175 303C165.7 312.4 165.7 327.6 175 336.1C184.4 346.3 199.6 346.3 208.1 336.1L255.1 289.9L303 336.1C312.4 346.3 327.6 346.3 336.1 336.1C346.3 327.6 346.3 312.4 336.1 303L289.9 255.1L336.1 208.1C346.3 199.6 346.3 184.4 336.1 175C327.6 165.7 312.4 165.7 303 175L255.1 222.1L208.1 175C199.6 165.7 184.4 165.7 175 175C165.7 184.4 165.7 199.6 175 208.1V208.1z"/></svg>
                        {:else if item.data.data.status === 'Running' || item.data.data.status === 'pending'}
                            <svg fill="blue" x="-9.5rem" y="-0.62rem" height="1.25rem" width="11rem" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--! Font Awesome Pro 6.1.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc. --><path d="M304 48C304 74.51 282.5 96 256 96C229.5 96 208 74.51 208 48C208 21.49 229.5 0 256 0C282.5 0 304 21.49 304 48zM304 464C304 490.5 282.5 512 256 512C229.5 512 208 490.5 208 464C208 437.5 229.5 416 256 416C282.5 416 304 437.5 304 464zM0 256C0 229.5 21.49 208 48 208C74.51 208 96 229.5 96 256C96 282.5 74.51 304 48 304C21.49 304 0 282.5 0 256zM512 256C512 282.5 490.5 304 464 304C437.5 304 416 282.5 416 256C416 229.5 437.5 208 464 208C490.5 208 512 229.5 512 256zM74.98 437C56.23 418.3 56.23 387.9 74.98 369.1C93.73 350.4 124.1 350.4 142.9 369.1C161.6 387.9 161.6 418.3 142.9 437C124.1 455.8 93.73 455.8 74.98 437V437zM142.9 142.9C124.1 161.6 93.73 161.6 74.98 142.9C56.24 124.1 56.24 93.73 74.98 74.98C93.73 56.23 124.1 56.23 142.9 74.98C161.6 93.73 161.6 124.1 142.9 142.9zM369.1 369.1C387.9 350.4 418.3 350.4 437 369.1C455.8 387.9 455.8 418.3 437 437C418.3 455.8 387.9 455.8 369.1 437C350.4 418.3 350.4 387.9 369.1 369.1V369.1z"/></svg>
                        {:else}
                            <svg fill="green" x="-9.5rem" y="-0.62rem" height="1.25rem" width="11rem" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--! Font Awesome Pro 6.1.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc. --><path d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM371.8 211.8C382.7 200.9 382.7 183.1 371.8 172.2C360.9 161.3 343.1 161.3 332.2 172.2L224 280.4L179.8 236.2C168.9 225.3 151.1 225.3 140.2 236.2C129.3 247.1 129.3 264.9 140.2 275.8L204.2 339.8C215.1 350.7 232.9 350.7 243.8 339.8L371.8 211.8z"/></svg>
                        {/if}
                        <text fill="#232323" x="-3rem" y="0.33rem" style="font-size: 1rem;">{item.data.data.name}</text>
                    </g>
                {/each}
            </g>
        </g>
    </svg>
</main>

