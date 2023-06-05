import { createRouter, defineEventHandler, useBase } from 'h3'
import {CurrentPollutionArgs, create_pollution_tile} from '~/server/services/polmap';
import { send } from 'h3'

const router = createRouter()

router.get('/:organisation_id/current/:x/:y/:zoom/tile.png', defineEventHandler(
    async (event) => {
        try {
            const tile_stream = await create_pollution_tile({
                organisation_id: event.context.params.organisation_id,
                x: parseInt(event.context.params.x),
                y: parseInt(event.context.params.y),
                zoom: parseInt(event.context.params.zoom),
            } as CurrentPollutionArgs);
            
            event.node.res.setHeader('Content-Type', 'image/png')
            return sendStream(event, tile_stream)
        } catch(error) {
            console.log(error)
        }

    }
))

export default useBase('/api/polmap', router.handler)