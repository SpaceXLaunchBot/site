import React from 'react';
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@material-ui/core';
import '../css/Commands.scss';

export default function Commands() {
  return (
    <TableContainer>
      <Table className="commandsTable">
        <TableHead>
          <TableRow>
            <TableCell>Command</TableCell>
            <TableCell>Description</TableCell>
            <TableCell>Required Permissions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          <TableRow>
            <TableCell>slb nextlaunch</TableCell>
            <TableCell>Send the latest launch schedule message to the current channel</TableCell>
            <TableCell>None</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>slb launch [number]</TableCell>
            <TableCell>
              Send the launch schedule message for the given launch number to the current channel
            </TableCell>
            <TableCell>None</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>slb add [type] [mentions]</TableCell>
            <TableCell>
              Add the current channel to the notification service with the given notification type
              (all, schedule, or launch). If you chose all or launch, the second part can be a list
              roles / channels / users to ping when a launch notification is sent
            </TableCell>
            <TableCell>Admin</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>slb remove</TableCell>
            <TableCell>Send the latest launch schedule message to the current channel</TableCell>
            <TableCell>Admin</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>slb info</TableCell>
            <TableCell>Send information about the bot to the current channel</TableCell>
            <TableCell>None</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>slb help</TableCell>
            <TableCell>List these commands</TableCell>
            <TableCell>None</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </TableContainer>
  );
}
