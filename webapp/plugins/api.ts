import { FeatureRepositoryFactory, IFeatureRepository } from "../repository/feature"
import { IReportRepository, ReportRepositoryFactory } from "~/repository/report";
import { defineNuxtPlugin } from '#app';
import { IMonitoringRepository, MonitoringRepositoryFactory } from "~/repository/monitoring";
import { IOrganisationRepository, OrganisationRepositoryFactory } from "~/repository/organisation";

interface IApiService {
    features: IFeatureRepository,
    reports: IReportRepository,
    monitoring: IMonitoringRepository,
    organisations: IOrganisationRepository
}

export default defineNuxtPlugin((nuxtApp) => {
    return {
        provide: {
            api: {
                features: FeatureRepositoryFactory(null),
                reports: ReportRepositoryFactory(null),
                monitoring: MonitoringRepositoryFactory(null),
                organisation: OrganisationRepositoryFactory(null)
            }
        }
    }
})