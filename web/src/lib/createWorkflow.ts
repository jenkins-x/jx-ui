export interface Data {
    pipeline: string;
    status: string;
    steps: step[],
}

interface step {
    kind: string;
    stage: {
        name: string;
        status: string;
        steps: subStep[]
    },
}

interface subStep {
        name: string;
        status: string;
}

export const createData = ( data: Data, d ) => {
    data.steps.map((step) => {
        if (step.kind.toLowerCase() === "stage" ){
            d.data = {
                name: step[`${step.kind.toLowerCase()}`].name,
                status: step[`${step.kind.toLowerCase()}`].status,
            }
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
    //                     children: [...]
    //                 }
    //             ]
    //         }
    //     ]
    // }

    for(let i = 0; i < data.steps.length; i++) {
        // main steps
        let step = data.steps[i];

        if (step.kind.toLowerCase() === "stage" ){
            // using t as a temporary object, for each step
            let t = {
                data: {},
                children: []
            };

            for(let j = step[`${step.kind.toLowerCase()}`].steps.length - 1; j >= 0; j--) {
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
    return d;
}
