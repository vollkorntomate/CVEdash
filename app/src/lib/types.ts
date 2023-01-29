export class CVESummary {
    id: string;
    desription: string;
    published: string;
    cvssBaseSeverity: string;
    cvssBaseScore: number;

    constructor(id: string,
        desription: string,
        published: string,
        cvssBaseSeverity: string,
        cvssBaseScore: number) {
            this.id = id
            this.desription = desription
            this.published = published
            this.cvssBaseSeverity = cvssBaseSeverity
            this.cvssBaseScore = cvssBaseScore
        }
};