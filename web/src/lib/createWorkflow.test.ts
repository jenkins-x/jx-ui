import { it, expect } from 'vitest'
import { createData, type Data } from './createWorkflow'

it('returns an object with fields: data (object) and children (array of one element with children.children.length = 0)', () => {
    const data: Data = {
        pipeline: "rajatgupta24/cluster-005/main",
        status: "Succeeded",
        steps: [
            {
                kind: "Stage",
                stage: {
                    name: "from build pack",
                    status: "Succeeded",
                    steps: [
                        {
                            name: "Admin Log",
                            status: "Succeeded",
                        },
                    ]
                }
            }
        ],
    }

    let d = {
        data: {
            name: String,
            status: String,
        },
        children: [],
    };

    // This is the expected result returned by createData
    //
    // const expectedResult = {
    //     data: {
    //         name: "from build pack",
    //         status: "Succeeded",
    //     },
    //     children: [
    //         {
    //             data: {
    //                 name: "Admin Log",
    //                 status: "Succeeded",
    //             },
    //             children: [],    
    //         },
    //     ]
    // }

    d = createData(data, d)
    expect(d).not.toContain([]);
    expect(d.data.name).toContain("from build pack");
    expect(d.children.length).toBe(1);
    expect(d.children[0].children.length).toBe(0);
    expect(d.children[0].data.name).toBe("Admin Log");
})

it('returns an object with fields: data (object) and children (array of one element with children.children.length = 1) ', () => {
    const data: Data = {
        pipeline: "rajatgupta24/cluster-005/main",
        status: "Succeeded",
        steps: [
            {
                kind: "Stage",
                stage: {
                    name: "from build pack",
                    status: "Succeeded",
                    steps: [
                        {
                            name: "Admin Log",
                            status: "Succeeded",
                        },
                        {
                            name: "Git Clone",
                            status: "Succeeded",
                        }
                    ]
                }
            }
        ],
    }

    let d = {
        data: {
            name: String,
            status: String,
        },
        children: [],
    };

    // This is the expected result returned by createData
    //
    // const expectedResult = {
    //     data: {
    //         name: "from build pack",
    //         status: "Succeeded",
    //     },
    //     children: [
    //         {
    //             data: {
    //                 name: "Admin Log",
    //                 status: "Succeeded",
    //             },
    //             children: [
    //                 {
    //                     data: {
    //                         name: "Git Clone",
    //                         status: "Succeeded",
    //                     },
    //                     children: []                        
    //                 },
    //             ],    
    //         },
    //     ]
    // }

    d = createData(data, d)
    expect(d).not.toContain([]);
    expect(d.data.name).toContain("from build pack");
    expect(d.children.length).toBe(1);
    expect(d.children[0].children.length).toBe(1);
    expect(d.children[0].data.name).toBe("Admin Log");
    expect(d.children[0].children[0].data.name).toBe("Git Clone");
})

it('returns an object with children array with only element where kind is stage) ', () => {
    const data: Data = {
        pipeline: "rajatgupta24/cluster-005/main",
        status: "Succeeded",
        steps: [
            {
                kind: "Stage",
                stage: {
                    name: "from build pack",
                    status: "Succeeded",
                    steps: [
                        {
                            name: "Admin Log",
                            status: "Succeeded",
                        },
                        {
                            name: "Git Clone",
                            status: "Succeeded",
                        }
                    ]
                }
            },
            {
                kind: "promote",
                stage: {
                    name: "from build pack",
                    status: "Succeeded",
                    steps: [
                        {
                            name: "Admin Log",
                            status: "Succeeded",
                        },
                    ]
                }
            }
        ],
    }

    let d = {
        data: {
            name: String,
            status: String,
        },
        children: [],
    };

    // This is the expected result returned by createData
    //
    // const expectedResult = {
    //     data: {
    //         name: "from build pack",
    //         status: "Succeeded",
    //     },
    //     children: [
    //         {
    //             data: {
    //                 name: "Admin Log",
    //                 status: "Succeeded",
    //             },
    //             children: [
    //                 {
    //                     data: {
    //                         name: "Git Clone",
    //                         status: "Succeeded",
    //                     },
    //                     children: []                        
    //                 },
    //             ],    
    //         },
    //     ]
    // }

    d = createData(data, d)
    expect(d.children.length).toBe(1);
})
