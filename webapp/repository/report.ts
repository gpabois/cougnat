import {booleanWithin, Feature, FeatureCollection, point} from '@turf/turf'
import local from './local'
import itertools from '../itertools'
import SnowflakeId from 'snowflake-id'

const snowflake = new SnowflakeId({
    mid : 42,
    offset : (2019-1970)*31536000*1000
});

interface ReportType {
    name: string,
    nature: string,
    label: string
}

interface Report {
    id: string | undefined,
    type: ReportType,
    owner: string | undefined,
    reported_at: Date | undefined,
    location: Feature
}

interface IReportRepository {
    GetTypes(): Promise<Array<ReportType>>
    Create(report: Report): Promise<string>
    GetMine(): Promise<Array<Report>>
}

const REPORT_TYPES_FIXTURES: Array<ReportType> = [
    {
        name: 'bitume',
        nature: 'smell',
        label: 'Odeur de bitume'
    }, {
        name: 'concasseur',
        nature: 'noise',
        label: 'Bruit de concasseur'
    }
];

class MockReportRepository {
    async GetTypes(): Promise<Array<ReportType>> {
        return Promise.resolve(REPORT_TYPES_FIXTURES)
    }

    async Create(report: Report): Promise<string> {
        report.id = snowflake.generate()
        report.reported_at = new Date()
        return await local.put('reports', report.id!, unref(toRaw(report)))
    }

    async GetMine(): Promise<Array<Report>> {
        return await itertools.async.toArray<Report>(local.cursor('reports'))
    }
}

function ReportRepositoryFactory(container: any): IReportRepository {
    return new MockReportRepository()
}

export {
    ReportType,
    Report,
    IReportRepository,
    ReportRepositoryFactory
}

