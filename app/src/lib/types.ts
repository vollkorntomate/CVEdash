export class CVESummary {
    id: string;
    desription: string;
    published: string;
    cvssSeverity: string;
    cvssScore: number;

    constructor(id: string,
        desription: string,
        published: string,
        cvssSeverity: string,
        cvssScore: number) {
            this.id = id
            this.desription = desription
            this.published = published
            this.cvssSeverity = cvssSeverity
            this.cvssScore = cvssScore
        }
};