import { FeatureRepositoryFactory, IFeatureRepository } from "../repository/feature"
import { IReportRepository, ReportRepositoryFactory } from "~/repository/report";
import { defineNuxtPlugin } from '#app';

interface IApiService {
    features: IFeatureRepository,
    reports: IReportRepository
}

export default defineNuxtPlugin((nuxtApp) => {
    return {
        provide: {
            api: {
                features: FeatureRepositoryFactory(null),
                reports: ReportRepositoryFactory(null)
            }
        }
    }
})